package models

type User struct {
	ID       uint   `json:"ID"`
	Nombre   string `json:"name"`
	Username string `json:"username"`
	IDPuesto uint   `json:"id_puesto"`
}
