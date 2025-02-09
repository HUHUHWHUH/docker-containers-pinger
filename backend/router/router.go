package router

import (
	"backend/handlers"
	"backend/services"
	"github.com/rs/cors"
	"net/http"
)

// Router - это кастомный маршрутизатор
type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

// addRoute добавляет маршрут в кастомный маршрутизатор
func (router *Router) addRoute(method string, path string, handler http.HandlerFunc) {
	if router.routes[path] == nil {
		router.routes[path] = make(map[string]http.HandlerFunc)
	}
	router.routes[path][method] = handler
}

func (router *Router) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	pathRoutes, exists := router.routes[request.URL.Path]
	if !exists {
		http.Error(responseWriter, "404 not found", http.StatusNotFound)
		return
	}

	handler, methodExists := pathRoutes[request.Method]
	if !methodExists {
		http.Error(responseWriter, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	handler(responseWriter, request)
}

// NewRouter создает кастомный роутер
func NewRouter(service *services.DockerService) *Router {
	router := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	router.addRoute("GET", "/containers", handlers.GetContainers(service))
	router.addRoute("POST", "/containers", handlers.UpsertContainer(service))
	router.addRoute("DELETE", "/containers", handlers.DeleteContainers(service))

	return router
}

// SetCORS устанавливает настройки CORS
func SetCORS(handler http.Handler) http.Handler {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9090"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})
	return corsHandler.Handler(handler)
}
