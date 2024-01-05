package shared_model

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	jwt.RegisteredClaims
	AuthorId string `json:"authorId"`
	Role     string `json:"role"`
}
