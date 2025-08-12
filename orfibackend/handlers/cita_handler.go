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
	// Validar que sea POST
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Método no permitido",
		})
		return
	}

	// Decodificar datos
	var cita Cita
	if err := json.NewDecoder(r.Body).Decode(&cita); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Datos inválidos",
		})
		return
	}

	// Guardar en BD
	db := config.GetDB()

	// Comprobar base de datos actual
	var currentDB string
	db.QueryRow("SELECT DATABASE()").Scan(&currentDB)
	fmt.Println("📌 Conectado a la base de datos:", currentDB)

	fmt.Printf("Datos recibidos: %+v\n", cita)

	err := db.QueryRow("SELECT DATABASE()").Scan(&currentDB)
	if err != nil {
		fmt.Println("❌ Error consultando base de datos actual:", err)
	} else {
		fmt.Println("📌 Conectado a la base de datos:", currentDB)
	}

	res, err := db.Exec(`INSERT INTO citas (nombre, servicio, fecha, telefono) VALUES (?, ?, ?, ?)`,
		cita.Nombre, cita.Servicio, cita.Fecha, cita.Telefono)

	if err != nil {
		fmt.Println("❌ Error en INSERT:", err)
	} else {
		rows, _ := res.RowsAffected()
		fmt.Println("✅ Filas insertadas:", rows)
	}

	// Notificación vía WhatsApp
	notifier, _ := factory.GetNotificationChannel("whatsapp")
	msg := fmt.Sprintf("Hola %s, tu cita de %s fue agendada para %s",
		cita.Nombre, cita.Servicio, cita.Fecha)
	notifier.Send(cita.Telefono, msg)

	// Respuesta exitosa en JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reserva creada con éxito",
	})
}
