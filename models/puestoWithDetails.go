package models

type PuestoWithDetails struct {
	ID            uint
	Nombre        string
	Descripcion   string
	Repetible     bool
	IDJefe        uint
	NombreFuncion string
}
