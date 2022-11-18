package middleware

import (
	"net/http"

	"github.com/jacobd39/edteam/go_api/authorization"
	"github.com/labstack/echo"
)

// Log .
//func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		log.Printf("Petición %q, método: %q", r.URL.Path, r.Method)
//		f(w, r)
//	}
//}

// Authentication .
func Authentication(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		_, err := authorization.ValidateToken(token)

		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "No tiene autorización"})
		}

		return f(c)
	}
}
