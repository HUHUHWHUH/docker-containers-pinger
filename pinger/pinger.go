package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"log"
	"os/exec"
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

	// Делаем пинг контейнера по IP адресу
	pingSuccess := pingContainerByIP(ip)
	container := t.DockerContainer{
		ContainerId: containerID,
		Ip:          ip,
		PingTime:    time.Now(),
		//LastSuccessfulPingTryTime: time.Now(),
		Status: inspectedCntr.State.Status,
		Name:   inspectedCntr.Name,
	}

	// Если пинг успешен, время последней успешной попытки обновляется
	if pingSuccess {
		container.LastSuccessfulPingTryTime = time.Now()
	}

	return container, nil

	//container := t.DockerContainer{
	//	ContainerId:               containerID,
	//	Ip:                        ip,
	//	PingTime:                  time.Now(),
	//	LastSuccessfulPingTryTime: time.Now(),
	//	Status:                    inspectedCntr.State.Status,
	//	Name:                      inspectedCntr.Name,
	//}
	//return container, nil
}

// pingContainerByIP пингует контейнер по IP
func pingContainerByIP(ip string) bool {
	cmd := exec.Command("ping", "-c", "1", "-w", "2", ip) // -c 1 - 1 запрос, -w 2 - тайм-аут в 2 секунды
	err := cmd.Run()
	if err != nil {
		log.Printf("Ошибка пинга: %v", err)
		return false
	}

	return true
}
