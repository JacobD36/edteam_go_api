package model

import "github.com/dgrijalva/jwt-go"

// Login .
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim contendrá el cuerpo que tendrá el token
type Claim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
