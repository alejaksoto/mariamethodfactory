package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification/config"
	"notification/factory"
)

type Cita struct {
	Nombre   string `json:"nombre"`
	Servicio string `json:"servicio"`
	Fecha    string `json:"fecha"`
	Telefono string `json:"telefono"`
}

func CrearCita(w http.ResponseWriter, r *http.Request) {
	var cita Cita
	if err := json.NewDecoder(r.Body).Decode(&cita); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	db := config.GetDB()
	_, err := db.Exec("INSERT INTO citas (nombre, servicio, fecha, telefono) VALUES (?, ?, ?, ?)",
		cita.Nombre, cita.Servicio, cita.Fecha, cita.Telefono)
	if err != nil {
		http.Error(w, "Error guardando en la base de datos", http.StatusInternalServerError)
		return
	}

	// Notificación vía WhatsApp (puedes cambiar a email o sms)
	notifier, _ := factory.GetNotificationChannel("whatsapp")
	msg := fmt.Sprintf("Hola %s, tu cita de %s fue agendada para %s", cita.Nombre, cita.Servicio, cita.Fecha)
	notifier.Send(cita.Telefono, msg)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Reserva creada con éxito"}`))
}
