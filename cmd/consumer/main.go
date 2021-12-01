package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"client.id": "goapp-consumer",
		"group.id": "goapp-group", // grupo de toda a aplicação, s/ esse group.id não iremos consumir a mensagem
		"auto.offset.reset": "earliest", // pega as mensagens desde o início.

	}
	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("error consumer", err.Error())
	}

	topics := []string{"teste"}
	c.SubscribeTopics(topics, nil)

	//ler as mensagens
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
		}
	}
}