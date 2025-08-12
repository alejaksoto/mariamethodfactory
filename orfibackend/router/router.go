package router

import (
	"net/http"
	"notification/handlers"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/citas", handlers.CrearCita)
	// registrar más rutas...
	return mux
}
