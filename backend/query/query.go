package query

import (
	"database/sql"
	"log"
)

// ShutdownUpdateStatuses обновляет все статусы контейнеров в таблице на 'shutdown'
func ShutdownUpdateStatuses(db *sql.DB) {
	query := "UPDATE dockerContainers SET status = 'shutdown'"
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Ошибка обновления статусов: %v", err)
	} else {
		log.Println("Все статусы обновлены на 'shutdown'")
	}
}

// CreateContainersTable создает таблицу dockerContainers, если она не существует.
func CreateContainersTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS dockerContainers(
    	id SERIAL PRIMARY KEY,
    	Container_name VARCHAR(255),
    	Container_id VARCHAR(255) UNIQUE NOT NULL,
    	IP VARCHAR(255),
    	PingTime TIMESTAMP,
    	LastSuccessfulPingTryTime TIMESTAMP,
        status VARCHAR(10)
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу: %v", err)
	}
}
