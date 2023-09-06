package models

import "gorm.io/gorm"

type Puesto struct {
	gorm.Model
	Nombre    string `json:"nombre"      form:"nombre"      binding:"required"  gorm:"unique"`
	IDJefe    uint   `json:"id_jefe"     form:"id_jefe" `
	IDUsuario uint   `json:"id_user"     form:"id_user" `
	IDFuncion uint   `json:"id_funcion"  form:"id_funcion"  binding:"required"`
}
