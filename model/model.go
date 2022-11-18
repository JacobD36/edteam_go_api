package model

import "errors"

var (
	//ErrPersonCanNotBeNil es el error que se lanza cuando la persona es nil
	ErrPersonCantNotBeNil = errors.New("La persona no puede ser nula")
	//ErrIDPersonDoesNotExists es el error que se lanza cuando el ID de la persona no existe
	ErrIDPersonDoesNotExists = errors.New("La persona no existe")
)
