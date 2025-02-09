package kafka

import (
	"backend/services"
	t "backend/types"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

// HandleKafkaMessages обрабатывает полученные сообщения из кафки
func HandleKafkaMessages(reader *kafka.Reader, cntrService *services.DockerService) {
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Ошибка при чтении сообщения из Kafka: %v", err)
			continue
		}

		// Декодируем сообщение в структуру DockerContainer
		var container t.DockerContainer
		err = json.Unmarshal(msg.Value, &container)
		if err != nil {
			log.Printf("Ошибка при декодировании сообщения: %v", err)
			continue
		}

		err = cntrService.UpsertContainer(container)
		if err != nil {
			log.Printf("Ошибка при сохранении контейнера: %v", err)
		}
	}
}

// CreateKafkaConsumer инициализирует потребителя сообщений
func CreateKafkaConsumer() *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9093"},
		Topic:   "docker-queue",
		GroupID: "backend-group",
	})
	return reader
}
