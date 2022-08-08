package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rhiadc/flow_test/config"
	ckafka "github.com/rhiadc/flow_test/kafka"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	ready chan bool
}

func main() {
	broker := "0.0.0.0:9092"
	topic := "topico-teste"
	config := config.New()
	config.Brokers = broker
	config.Topic = topic
	config.Group = "teste"

	producer := ckafka.NewProducer(broker, topic)
	message, _ := json.Marshal("TESTEEE3")
	msg := kafka.Message{Value: message}
	if err := producer.WriteMessages(context.Background(), msg); err != nil {
		log.Fatal(err)
	}

	consumer := ckafka.NewConsumer(config)
	if err := consumer.SyncConsume(context.Background(), 3); err != nil {
		log.Println("Exited with error: ", err)
	}
	log.Println("Successfully processed")

}

//todos
//buscar nome dos tópicos e brokers de money-moved-approved, firt-transaction, score
//conseguir produzir e ler, ex: produzido no money-moved-approved deve refletir no first-transaction
//implementar casos descritos (unitarios e ponta a ponta)
//estudar onde subir esse código/como rodar

func TestCasesFromMoneyMovedApproved() {
	//StartProducerAndConsumer()
	//create msg cases
	//CreateComplexEvent()
	//consume
	//assert returns
	//CloseProducerAndConsumer()
}

func TestCasesFromFirstTransaction() {
	//StartProducerAndConsumer()
	//create msg cases
	//SendMessageToMoneyMovedApprovedTopic
	//consume
	//assert returns
	//CloseProducerAndConsumer()
}
func TestCasesFromScore() {
	//StartProducerAndConsumer()
	//create msg cases
	//SendMessageToFirstTransaction()
	//consume
	//assert returns
	//CloseProducerAndConsumer()
}

func TestCasesFromComplexEventsToConnector() {
	//startProducerAndConsumer()
	//create msg cases
	//send message to adapter
	//consume from score
	//assert score
	//CloseProducerAndCOnsumer()
}
func StartProducerAndConsumer()             {}
func CloseProducerAndConsumer()             {}
func CreateComplexEvent()                   {}
func SendMessageToMoneyMovedApprovedTopic() {}
func SendMessageToFirstTransaction()        {}
