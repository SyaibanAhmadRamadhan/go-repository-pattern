package dto

import "github.com/dgrijalva/jwt-go"

type TokenClaim struct {
	UserId   string `json:"userr_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.StandardClaims
}
