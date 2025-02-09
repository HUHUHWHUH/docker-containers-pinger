package types

import "time"

// DockerContainer - структура контейнера
type DockerContainer struct {
	ID                        int
	Name                      string    `json:"container_name"`
	ContainerId               string    `json:"container_id"`
	Ip                        string    `json:"ip"`
	PingTime                  time.Time `json:"pingTime"`
	LastSuccessfulPingTryTime time.Time `json:"lastSuccessfulPingTryTime"`
	Status                    string    `json:"Status"`
}
