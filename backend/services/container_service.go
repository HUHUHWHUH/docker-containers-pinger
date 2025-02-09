package services

import (
	t "backend/types"
	"database/sql"
	"fmt"
	"time"
)

// DockerService управляет контейнерами в памяти
type DockerService struct {
	db *sql.DB
}

// NewDockerService создает новый докер сервис
func NewDockerService(db *sql.DB) *DockerService {
	return &DockerService{db}
}

// UpsertContainer вставка(если контейнера нет в бд) или обновление контейнера в бд
func (dockerService *DockerService) UpsertContainer(cntr t.DockerContainer) error {
	name := cntr.Name
	containerId := cntr.ContainerId
	ip := cntr.Ip
	pingTime := cntr.PingTime
	lastSuccessfulTryTime := cntr.LastSuccessfulPingTryTime
	status := cntr.Status

	if pingTime.IsZero() {
		pingTime = time.Now()
	}
	if lastSuccessfulTryTime.IsZero() {
		lastSuccessfulTryTime = time.Now()
	}

	query := `
		INSERT INTO dockerContainers (Container_name, Container_id, IP, PingTime, LastSuccessfulPingTryTime, Status)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (Container_id)
		DO UPDATE SET Container_name = EXCLUDED.Container_name, IP = EXCLUDED.IP, PingTime = EXCLUDED.PingTime, LastSuccessfulPingTryTime = EXCLUDED.LastSuccessfulPingTryTime, Status = EXCLUDED.Status;
	`

	_, err := dockerService.db.Exec(query, name, containerId, ip, pingTime, lastSuccessfulTryTime, status)
	if err != nil {
		return fmt.Errorf("не удалось вставить или обновить данные контейнера с id %s: %v", containerId, err)
	}
	return nil
}

// GetAllCntrs возвращает все контейнеры в бд
func (dockerService *DockerService) GetAllCntrs() ([]t.DockerContainer, error) {
	query := "SELECT id, container_name, container_id, ip, pingTime, LastSuccessfulPingTryTime, Status FROM dockerContainers"
	rows, err := dockerService.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить контейнеры: %v", err)
	}
	defer rows.Close()
	containers := []t.DockerContainer{}
	for rows.Next() {
		var cntr t.DockerContainer
		err = rows.Scan(&cntr.ID, &cntr.Name, &cntr.ContainerId, &cntr.Ip, &cntr.PingTime, &cntr.LastSuccessfulPingTryTime, &cntr.Status)
		if err != nil {
			return nil, fmt.Errorf("ошибка при получении контейнеров: %v", err)
		}
		containers = append(containers, cntr)
	}
	if err = rows.Err(); err != nil {
		return containers, fmt.Errorf("ошибка при чтении строк: %v", err)
	}

	return containers, nil
}

// DeleteShutdownContainers удаляет все контейнеры состоянием 'shutdown'
func (dockerService *DockerService) DeleteShutdownContainers() error {
	query := "DELETE FROM dockerContainers WHERE status = 'shutdown'"
	_, err := dockerService.db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при удалении контейнеров: %v", err)
	}
	return nil
}
