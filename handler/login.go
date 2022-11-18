package handler

import (
	"net/http"

	"github.com/jacobd39/edteam/go_api/authorization"
	"github.com/jacobd39/edteam/go_api/model"
	"github.com/labstack/echo"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(c echo.Context) error {
	data := model.Login{}

	err := c.Bind(&data)

	if err != nil {
		response := newResponse(Error, "estructura no válida", nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "credenciales no válidas", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	token, err := authorization.GenerateToken(&data)

	if err != nil {
		response := newResponse(Error, "no se pudo generar el token", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	dataToken := map[string]string{"token": token}

	resp := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "contacto@gmail.com" && data.Password == "123456"
}
