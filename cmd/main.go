package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jacobd39/edteam/go_api/authorization"
	"github.com/jacobd39/edteam/go_api/handler"
	"github.com/jacobd39/edteam/go_api/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")

	if err != nil {
		log.Fatalf("no se pudieron cargar los archivos de certificados: %v", err)
	}

	store := storage.NewMemory()

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RouteLogin(e, &store)
	handler.RoutePerson(e, &store)

	log.Println("Servidor iniciado...")

	err = e.Start(":8080")

	if err != nil {
		log.Printf("Error en el servidor: %v", err)
	}
}
