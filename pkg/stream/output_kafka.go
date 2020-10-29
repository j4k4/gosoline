package stream

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type KafkaProducer interface {
	Close()
	Events() chan kafka.Event
	Flush(timeoutMs int) int
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

type KafkaOutputSettings struct {
	Topic        string        `cfg:"topic" validate:"required"`
	FlushTimeout time.Duration `cfg:"flush_timeout" default:"3s" validate:"required"`
}

type kafkaOutput struct {
	producer KafkaProducer
	settings *KafkaOutputSettings
}

func NewKafkaOutput(logger mon.Logger, settings *KafkaOutputSettings) *kafkaOutput {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	producer.Events()
	if err != nil {
		logger.Fatalf(err, "can not create kafka producer for output")
	}

	return NewKafkaOutputWithInterfaces(producer, settings)
}

func NewKafkaOutputWithInterfaces(producer KafkaProducer, settings *KafkaOutputSettings) *kafkaOutput {
	return &kafkaOutput{
		producer: producer,
		settings: settings,
	}
}

func (k *kafkaOutput) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		k.producer.Close()
	}()

	var event kafka.Event
	var remaining int

	for event = range k.producer.Events() {
		switch ev := event.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
			} else {
				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
			}
		}
	}

	remaining = k.producer.Flush(int(k.settings.FlushTimeout.Milliseconds()))

	if remaining > 0 {
		return fmt.Errorf("wasn't able to flush all remaining messages")
	}

	return nil
}

func (k *kafkaOutput) WriteOne(ctx context.Context, msg WritableMessage) error {
	return k.Write(ctx, []WritableMessage{msg})
}

func (k *kafkaOutput) Write(_ context.Context, batch []WritableMessage) error {
	var err error
	var bytes []byte

	for _, msg := range batch {
		if bytes, err = msg.MarshalToBytes(); err != nil {
			return fmt.Errorf("can not marshal message to bytes: %w", err)
		}

		err = k.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &k.settings.Topic, Partition: kafka.PartitionAny},
			Value:          bytes,
		}, nil)
	}

	return nil
}
