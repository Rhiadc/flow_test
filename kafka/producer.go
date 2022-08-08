package kafka

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	kafkaWriter *kafka.Writer
}

func NewProducer(brokers, topic string) *Producer {

	return &Producer{
		kafkaWriter: kafka.NewWriter(
			kafka.WriterConfig{
				Brokers: strings.Split(brokers, ","),
				Topic:   topic,
			},
		),
	}
}

func (l Producer) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return l.kafkaWriter.WriteMessages(ctx, msgs...)
}

func (l Producer) Close() error {
	return l.kafkaWriter.Close()
}
