package handler

import "github.com/jacobd39/edteam/go_api/model"

//Storage es una interfaz que define los métodos que se pueden usar en el handler
type Storage interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}
