package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Data struct {
	Album  string
	Year   string
	Artist string
	Ranked string
}

func main() {

	conf := ReadConfig()
	fmt.Println(conf)
	conf["group.id"] = "mygroupid"
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	topic := "mytopic"
	c.SubscribeTopics([]string{topic}, nil)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			c.Close()
			return
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}
}

func ReadConfig() kafka.ConfigMap {

	m := make(map[string]kafka.ConfigValue)

	m["bootstrap.servers"] = "kafka-service:9092"
	return m

}
