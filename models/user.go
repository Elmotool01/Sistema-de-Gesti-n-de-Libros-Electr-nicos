package models

type User struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"`
}
