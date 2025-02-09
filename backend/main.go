package main

import (
	"backend/kafka"
	"backend/query"
	"backend/router"
	"backend/services"
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Не удалось открыть БД: %v", err)
	}
	defer db.Close()

	query.CreateContainersTable(db)

	cntrService := services.NewDockerService(db)

	reader := kafka.CreateKafkaConsumer()
	defer reader.Close()
	go kafka.HandleKafkaMessages(reader, cntrService) // Запускаем обработку сообщений из Kafka в отдельной горутине (через кафку происходит взаимодействие с пингером)

	customRouter := router.NewRouter(cntrService)
	settedCorsRouter := router.SetCORS(customRouter)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: settedCorsRouter,
	}

	quit := make(chan os.Signal, 1) // Канал для получения сигнала завершения
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине (сервер принимает запросы с фронтенда)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка при запуске сервера: %v", err)
		}
	}()

	// Реализация Graceful Shutdown
	<-quit // Ожидание сигнал завершения работы
	query.ShutdownUpdateStatuses(db)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %v", err)
	}

	log.Println("Сервер завершил работу")
}
