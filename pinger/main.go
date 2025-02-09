package main

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
	"pinger/kafka"
	"time"
)

func main() {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithVersion("1.46")) // указываем последнюю поддерживаемую версию
	if err != nil {
		log.Fatalf("Ошибка в создании клиента: %v", err)
	}
	defer dockerClient.Close()

	producer := kafka.CreateKafkaProducer()

	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		log.Fatalf("Ошибка в получении контейнеров: %v", err)
	}
	//apiURL := "http://backend:8000/containers"

	// Пингуем и отправляем результат каждые 10 секунд
	for {
		for _, cntr := range containers {
			pingResult, err := pingContainer(dockerClient, cntr.ID)
			if err != nil {
				log.Printf("Ошибка пинга контейнера %s: %v", cntr.ID, err)
			} else {
				resultJSON, err := json.Marshal(pingResult)
				if err != nil {
					log.Printf("Ошибка сериализации результата пинга: %v", err)
					continue
				}

				kafka.SendToKafka(producer, string(resultJSON))

				//err = sendPingResult(apiURL, pingResult)
				//if err != nil {
				//	log.Printf("Ошибка отправки результата пинга: %v", err)
				//}
			}
		}
		// Пауза 10 сек между пингами
		time.Sleep(10 * time.Second)
	}
}
