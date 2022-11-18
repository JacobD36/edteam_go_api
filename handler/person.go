package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/jacobd39/edteam/go_api/model"
	"github.com/labstack/echo"
)

type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(c echo.Context) error {
	data := model.Person{}

	err := c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "Error al leer el body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Create(&data)

	if err != nil {
		response := newResponse(Error, "Ocurrió un problema al intentar crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona creada correctamente", nil)
	return c.JSON(http.StatusCreated, response)
}

func (p *person) getAll(c echo.Context) error {
	persons, err := p.storage.GetAll()

	if err != nil {
		response := newResponse(Error, "Ocurrió un problema al intentar obtener las personas", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", persons)
	return c.JSON(http.StatusOK, response)
}

func (p *person) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "El ID debe ser de tipo entero positivo", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := model.Person{}
	err = c.Bind(&data)
	if err != nil {
		response := newResponse(Error, "Error al leer el body", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Update(id, &data)
	if err != nil {
		response := newResponse(Error, "Ocurrió un problema al intentar actualizar la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona actualizada correctamente", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response := newResponse(Error, "El ID debe ser de tipo entero positivo", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	err = p.storage.Delete(id)
	if err != nil {
		if errors.Is(err, model.ErrIDPersonDoesNotExists) {
			response := newResponse(Error, "El ID no existe", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := newResponse(Error, "Ocurrió un problema al intentar eliminar el registro de la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "Persona eliminada de forma correcta", nil)
	return c.JSON(http.StatusOK, response)
}

func (p *person) getByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		if err != nil {
			response := newResponse(Error, "El ID debe ser de tipo entero positivo", nil)
			c.JSON(http.StatusBadRequest, response)
		}
	}

	data, err := p.storage.GetByID(id)

	if err != nil {
		if errors.Is(err, model.ErrIDPersonDoesNotExists) {
			response := newResponse(Error, "El ID no existe", nil)
			return c.JSON(http.StatusBadRequest, response)
		}

		response := newResponse(Error, "Ocurrió un problema al intentar obtener el registro de la persona", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := newResponse(Message, "OK", data)
	return c.JSON(http.StatusOK, response)
}
