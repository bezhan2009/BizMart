package kafka

import (
	"BizMart/internal/app/models"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
	topic    string
}

var prdKafka KafkaProducer

func CreateProducer(kafkaParams models.KafkaParams) error {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", kafkaParams.Host, kafkaParams.Port),
	})
	if err != nil {
		return err
	}

	prdKafka.producer = kafkaProducer
	prdKafka.topic = kafkaParams.Topic

	return nil
}

func GetKafkaProducer() KafkaProducer {
	return prdKafka
}

func SendMessage(user models.User) error {
	// Кодируем в JSON
	messageBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}

	// Отправляем сообщение в Kafka
	err = prdKafka.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &prdKafka.topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}, nil)

	if err != nil {
		return err
	}

	fmt.Println("Message sent")
	return nil
}
