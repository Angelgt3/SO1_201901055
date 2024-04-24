package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
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
				team := rand.Intn(2) + 1
				err = rdb.HIncrBy(ctx, "teams", "team"+strconv.Itoa(team), 1).Err()
				if err != nil {
					panic(err)
				}
				err = rdb.Incr(ctx, "total").Err()
				if err != nil {
					panic(err)
				}
				fmt.Fprintf(os.Stderr, "%v\n", map[string]int{"team" + strconv.Itoa(team): 1})
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Mantener el programa en ejecución
	<-sigchan
}
