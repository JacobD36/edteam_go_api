package authorization

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jacobd39/edteam/go_api/model"
)

// GenerateToken genera un token para un usuario
func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "EDteam",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)

	signedToken, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateToken valida el token
func ValidateToken(t string) (model.Claim, error) {
	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyFunction)

	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("token no válido")
	}

	claim, ok := token.Claims.(*model.Claim)

	if !ok {
		return model.Claim{}, errors.New("no se pudo obtener el claim")
	}

	return *claim, nil
}

func verifyFunction(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
