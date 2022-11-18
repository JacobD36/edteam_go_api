package handler

import (
	"github.com/jacobd39/edteam/go_api/middleware"
	"github.com/labstack/echo"
)

// RoutePerson es la ruta para la persona
func RoutePerson(e *echo.Echo, storage Storage) {
	h := newPerson(storage)
	person := e.Group("/v1/persons")
	person.Use(middleware.Authentication)
	person.POST("/create", h.create)
	person.GET("/getAll", h.getAll)
	person.PUT("/update/:id", h.update)
	person.DELETE("/delete/:id", h.delete)
	person.GET("/get-by-id/:id", h.getByID)
	//mux.HandleFunc("/v1/persons/create", middleware.Authentication(middleware.Log(h.create)))
	//mux.HandleFunc("/v1/persons/getAll", middleware.Authentication(middleware.Log(h.getAll))) //endpoint alternativo get-all
	//mux.HandleFunc("/v1/persons/update", middleware.Authentication(middleware.Log(h.update)))
	//mux.HandleFunc("/v1/persons/delete", middleware.Authentication(middleware.Log(h.delete)))
}

// RouteLogin es la ruta para el login
func RouteLogin(e *echo.Echo, storage Storage) {
	h := newLogin(storage)
	e.POST("/v1/login", h.login)
}
