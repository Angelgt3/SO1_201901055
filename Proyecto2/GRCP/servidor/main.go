package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "servidor/protoServidor"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port        = ":3001"
	kafkaBroker = "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092"
)

var kafkaTopic = "mytopic"

func produceToKafka(data string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		return err
	}
	defer p.Close()

	deliveryChan := make(chan kafka.Event)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte(data),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}

	return nil
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetArtist())
	data := fmt.Sprintf("Year: %s, Album: %s, Artist: %s, Ranked: %s", in.GetYear(), in.GetAlbum(), in.GetArtist(), in.GetRanked())
	fmt.Println(data)

	err := produceToKafka(data)
	if err != nil {
		log.Printf("Failed to produce message to Kafka: %s", err)
	}

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el album"}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
