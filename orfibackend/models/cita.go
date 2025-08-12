package models

type Cita struct {
	Nombre   string `json:"nombre"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
	Hora     string `json:"hora"`
}
