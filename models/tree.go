package models

type Tree struct {
	Puesto       PuestoWithDetails
	Subordinados []Tree
}
