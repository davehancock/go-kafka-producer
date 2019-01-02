package messaging

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/daves125125/go-producer/config"
)

type kafkaProducer struct {
	*kafka.Producer
}

func (kp *kafkaProducer) PublishMessages(topic string, msgs []string) {

	for _, word := range msgs {
		err := kp.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)

		if err != nil {
			fmt.Printf("Produce failed: %v\n", err)
		}
	}

	// Wait for message deliveries before shutting down
	remainingMsgs := kp.Flush(10 * 1000)

	fmt.Println("Finished message producing. Messages Not sent:", remainingMsgs)
}

func (kp *kafkaProducer) Close() {
	kp.Producer.Close()
}

func NewKafkaProducer(cfg config.Provider) Producer {

	kafkaCfg := &kafka.ConfigMap{"bootstrap.servers": cfg.GetString("kafka.bootstrap.servers")}
	debugOpts := cfg.GetString("kafka.debug")
	if debugOpts != "" {
		kafkaCfg.SetKey("debug", debugOpts)
	}

	p, err := kafka.NewProducer(kafkaCfg)
	if err != nil {
		fmt.Println("Error creating producer", err)
		panic(err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			fmt.Println("Producer event: ", e)
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &kafkaProducer{p}
}
