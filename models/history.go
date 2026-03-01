package models

type History struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Accion string `json:"accion"`
	Fecha  string `json:"fecha"`
}
