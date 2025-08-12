package models

type Cita struct {
	Nombre   string `json:"nombre"`
	Servicio string `json:"servicio"`
	Fecha    string `json:"fecha"`
	Telefono string `json:"telefono"`
}
