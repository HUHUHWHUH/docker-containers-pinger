package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

// CreateKafkaProducer инициализирует отправителя
func CreateKafkaProducer() *kafka.Writer {
	producer := &kafka.Writer{
		Addr:     kafka.TCP("kafka:9093"),
		Topic:    "docker-queue",
		Balancer: &kafka.LeastBytes{},
	}
	return producer
}

// SendToKafka отправляет результата пинга в Kafka
func SendToKafka(writer *kafka.Writer, message string) {
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения в кафку: %s", err)
	} else {
		log.Println("Сообщение отправлено в кафку")
	}
}
