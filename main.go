package main

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github/cassiolpaixao/simulator-go/application/kafka"
	"github/cassiolpaixao/simulator-go/infra/kafka"
	"log"
)

func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main()  {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}

