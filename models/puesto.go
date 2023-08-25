package models

import "gorm.io/gorm"

type Puesto struct {
	gorm.Model
	Nombre      string `json:"nombre"      form:"nombre"      binding:"required"  gorm:"unique"`
	//Descripcion string `json:"descripcion" form:"descripcion" binding:"required"`
	//Repetible   bool   `json:"repetible"   form:"repetible"`
	IDJefe      uint   `json:"id_jefe"     form:"id_jefe" `
	IDFuncion   uint   `json:"id_funcion"  form:"id_funcion"  binding:"required"`
}
