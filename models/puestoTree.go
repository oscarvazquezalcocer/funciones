package models

type PuestoTree struct {
	Puesto       Puesto
	Subordinados []PuestoTree
}
