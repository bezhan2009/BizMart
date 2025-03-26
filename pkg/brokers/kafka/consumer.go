package kafka

import (
	"BizMart/internal/app/models"
	"BizMart/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConsumeUserMessages(kafkaParams models.KafkaParams) (models.User, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%d", kafkaParams.Host, kafkaParams.Port),
		"group.id":          kafkaParams.GroupID,
		"auto.offset.reset": kafkaParams.AutoOffsetReset,
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	err = consumer.Subscribe(kafkaParams.Topic, nil)
	if err != nil {
		return models.User{}, err
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			var user models.User
			if err = json.Unmarshal(msg.Value, &user); err != nil {
				logger.Error.Printf("[kafka.ConsumeUserMessages] Failed to unmarshal kafka messages: %s\n", err)

				return models.User{}, err
			} else {
				return user, nil
			}
		} else {
			logger.Error.Printf("[kafka.ConsumeUserMessages] Failed to read kafka messages: %s\n", err)

			return models.User{}, err
		}
	}
}
