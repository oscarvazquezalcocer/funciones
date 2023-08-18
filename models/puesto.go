package models

import "gorm.io/gorm"

type Puesto struct {
	gorm.Model
	Nombre      string `json:"nombre"      form:"nombre" gorm:"unique" binding:"required"`
	Descripcion string `json:"descripcion" form:"descripcion" binding:"required"`
	ID_Jefe     uint   `json:"id_jefe"   form:"id_jefe" `
}
