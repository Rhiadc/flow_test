package kafka

import (
	"context"
	"fmt"

	"github.com/rhiadc/flow_test/config"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	Reader *kafka.Reader
}

func NewConsumer(config *config.Config) *Consumer {
	return &Consumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.Brokers},
			Topic:   config.Topic,
			GroupID: config.Group,
		}),
	}
}

func (r *Consumer) SyncConsume(ctx context.Context, val int) error {
	for ref := 0; ref < val; ref++ {

		message, err := r.Reader.FetchMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		r.Reader.CommitMessages(ctx, message)
		fmt.Println("received: ", string(message.Value))
	}
	return r.Reader.Close()
}

func (r *Consumer) Close() error {
	return r.Reader.Close()
}
