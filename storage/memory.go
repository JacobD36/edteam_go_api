package storage

import (
	"fmt"

	"github.com/jacobd39/edteam/go_api/model"
)

// Memory es una estructura que guarda información en memoria
type Memory struct {
	currentID int
	Persons   map[int]model.Person
}

// NewMemory crea una nueva estructura de tipo Memory
func NewMemory() Memory {
	persons := make(map[int]model.Person)

	return Memory{
		currentID: 0,
		Persons:   persons,
	}
}

// Create crea una nueva persona
func (m *Memory) Create(person *model.Person) error {
	if person == nil {
		return model.ErrPersonCantNotBeNil
	}

	m.currentID++
	m.Persons[m.currentID] = *person

	return nil
}

// Update actualiza una persona en el slice de memoria
func (m *Memory) Update(ID int, person *model.Person) error {
	if person == nil {
		return model.ErrPersonCantNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	m.Persons[ID] = *person

	return nil
}

// Delete borra de la memoria la persona
func (m *Memory) Delete(ID int) error {
	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	delete(m.Persons, ID)

	return nil
}

// GetByID retorna una persona por el ID
func (m *Memory) GetByID(ID int) (model.Person, error) {
	person, ok := m.Persons[ID]
	if !ok {
		return person, fmt.Errorf("ID: %d: %w", ID, model.ErrIDPersonDoesNotExists)
	}

	return person, nil
}

// GetAll retorna todas las personas que están en la memoria
func (m *Memory) GetAll() (model.Persons, error) {
	var result model.Persons
	for _, v := range m.Persons {
		result = append(result, v)
	}

	return result, nil
}
