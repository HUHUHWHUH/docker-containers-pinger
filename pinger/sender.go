package main

// sendPingResult отправляет результат пинга на бэкенд для сохранения
//func sendPingResult(apiURL string, result t.DockerContainer) error {
//	// Сериализуем структуру DockerContainer в JSON
//	jsonData, err := json.Marshal(result)
//	if err != nil {
//		return fmt.Errorf("ошибка сериализации в JSON: %v", err)
//	}
//
//	// Отправляем POST-запрос на API
//	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
//	if err != nil {
//		return fmt.Errorf("ошибка отправки POST-запроса: %v", err)
//	}
//	defer resp.Body.Close() // закрываем тело запроса
//
//	// Проверка статус-кода ответа
//	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
//		return fmt.Errorf("неожиданный статус-код: %d", resp.StatusCode)
//	}
//
//	return nil
//}
