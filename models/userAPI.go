package models

type UserAPI struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
	IDPuesto uint   `json:"id_puesto"`
}
