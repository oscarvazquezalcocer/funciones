package models

type PuestoWithDetails struct {
	ID            uint
	Nombre        string
	Repetible     bool
	IDFuncion     uint
	IDJefe        uint
	NombreFuncion string
}
