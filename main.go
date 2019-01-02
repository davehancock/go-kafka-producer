package main

import (
	"fmt"

	"github.com/daves125125/go-producer/config"
	"github.com/daves125125/go-producer/messaging"
)

func main() {

	fmt.Println("Starting...")

	cfg := config.InitConfigProvider()

	p := messaging.NewKafkaProducer(cfg)
	defer p.Close()

	msgs := []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"}
	topic := cfg.GetString("kafka.producer.topic")

	p.PublishMessages(topic, msgs)
	p.PublishMessages(topic, msgs[0:3])

	fmt.Println("Tearing down...")
}
