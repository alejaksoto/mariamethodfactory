package main

import (
	"log"
	"net/http"
	"notification/router"
)

func main() {
	r := router.SetupRoutes()

	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		r.ServeHTTP(w, req)
	})

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
