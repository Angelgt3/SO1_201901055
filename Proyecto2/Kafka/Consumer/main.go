package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-redis/redis/v8"
)

type Data struct {
	Album  string `json:"Album"`
	Year   string `json:"Year"`
	Artist string `json:"Artist"`
	Ranked string `json:"Ranked"`
}

func main() {
	// Configuración de Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092",
		"group.id":          "mygroupid",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Failed to create Kafka consumer: %s", err)
		os.Exit(1)
	}

	// Configuración de Redis
	ctx := context.Background()
	rhost := os.Getenv("REDIS_HOST")
	rauth := os.Getenv("REDIS_AUTH")

	rdb := redis.NewClient(&redis.Options{
		Addr:     rhost + ":6379",
		Password: rauth,
		DB:       0,
	})

	topic := "mytopic"
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Failed to subscribe to topic: %s", err)
		os.Exit(1)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Función para consumir mensajes de Kafka y actualizar Redis
	go func() {
		for {
			select {
			case sig := <-sigchan:
				fmt.Printf("Caught signal %v: terminating\n", sig)
				c.Close()
				os.Exit(0)
			default:

				ev, err := c.ReadMessage(100 * time.Millisecond)
				if err != nil {
					continue
				}
				fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

				// Actualizar datos en Redis
				err = processAndUpdateRedis(ctx, rdb, string(ev.Value))
				if err != nil {
					fmt.Printf("Failed to process and update Redis: %s\n", err)
				}
			}
		}
	}()

	// Mantener el programa en ejecución
	<-sigchan
}

func processAndUpdateRedis(ctx context.Context, rdb *redis.Client, data string) error {
	// Procesar la cadena de datos para extraer los valores
	values := strings.Split(data, ", ")
	name := strings.Split(values[0], ": ")[1]
	album := strings.Split(values[1], ": ")[1]
	year := strings.Split(values[2], ": ")[1]

	// Clave del hash en Redis
	hashKey := fmt.Sprintf("%s:%s:%s", name, album, year)

	// Actualizar el contador de votos en Redis utilizando HINCRBY
	err := rdb.HIncrBy(ctx, hashKey, "votes", 1).Err()
	if err != nil {
		return err
	}

	fmt.Printf("Se agregó 1 voto para %s - %s (%s)\n", name, album, year)
	return nil
}
