package router

import (
	"net/http"
	"notification/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/api/citas", handlers.CrearCita)
}
