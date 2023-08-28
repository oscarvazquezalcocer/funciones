package models

type Tree struct {
	Puesto       Puesto
	User         UserAPI
	Subordinados []Tree
}
