package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	t "pinger/types"
	"time"
)

// pingContainer выполняет пинг контейнера и возвращает структуру DockerContainer
func pingContainer(dockerClient *client.Client, containerID string) (t.DockerContainer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	inspectedCntr, err := dockerClient.ContainerInspect(ctx, containerID)
	if err != nil {
		return t.DockerContainer{}, fmt.Errorf("ошибка при инспекции контейнера: %v", err)
	}

	// Получаем IP контейнера
	ip := ""
	if len(inspectedCntr.NetworkSettings.Networks) > 0 {
		// Перебираем сети, к которым подключен контейнер (хотя в этом проекте у каждого контейнера только одна сеть)
		for _, network := range inspectedCntr.NetworkSettings.Networks {
			ip = network.IPAddress
		}
	}

	container := t.DockerContainer{
		ContainerId:               containerID,
		Ip:                        ip,
		PingTime:                  time.Now(),
		LastSuccessfulPingTryTime: time.Now(),
		Status:                    inspectedCntr.State.Status,
		Name:                      inspectedCntr.Name,
	}

	return container, nil
}
