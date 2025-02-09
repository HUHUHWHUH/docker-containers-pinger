package handlers

import (
	"backend/services"
	t "backend/types"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContainers возвращает все контейнеры в бд
func GetContainers(service *services.DockerService) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		cntrs, err := service.GetAllCntrs()
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Ошибка получения контейнеров: %v", err), http.StatusInternalServerError)
			return
		}
		responseWriter.Header().Set("Content-Type", "application/json")
		if err = json.NewEncoder(responseWriter).Encode(cntrs); err != nil {
			http.Error(responseWriter, "Ошибка при преобразовании обьекта в JSON", http.StatusInternalServerError)
		}
	}
}

// UpsertContainer создает нового пользователя
func UpsertContainer(service *services.DockerService) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		var cntr t.DockerContainer
		if err := json.NewDecoder(request.Body).Decode(&cntr); err != nil {
			http.Error(responseWriter, "Некорректное тело запроса", http.StatusBadRequest)
			return
		}
		if err := service.UpsertContainer(cntr); err != nil {
			http.Error(responseWriter, fmt.Sprintf("Ошибка при попытке вставки/обновления контейнера в бд: %v", err), http.StatusInternalServerError)
			return
		}

		responseWriter.WriteHeader(http.StatusCreated)
		responseWriter.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(responseWriter).Encode(cntr); err != nil {
			http.Error(responseWriter, "Ошибка при преобразовании обьекта в JSON", http.StatusInternalServerError)
		}
	}
}

// DeleteContainers удаляет все контейнеры со статусом "shutdown"
func DeleteContainers(service *services.DockerService) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		err := service.DeleteShutdownContainers()
		if err != nil {
			http.Error(responseWriter, fmt.Sprintf("Ошибка при удалении контейнеров со статусом 'shutdown': %v", err), http.StatusInternalServerError)
			return
		}
		responseWriter.WriteHeader(http.StatusOK)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write([]byte("Контейнеры со статусом 'shutdown' успешно удалены"))
	}
}
