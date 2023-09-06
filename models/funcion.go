package models

import "gorm.io/gorm"

type Funcion struct {
	gorm.Model
	Nombre    string `json:"nombre"      form:"nombre"      binding:"required"  gorm:"unique"`
	IDJefe    uint   `json:"id_jefe"     form:"id_jefe" `
	IDUsuario uint   `json:"id_user"     form:"id_user" `
	IDPuesto  uint   `json:"id_puesto"   form:"id_puesto"  binding:"required"`
}
