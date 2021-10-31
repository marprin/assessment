package jwt

import "github.com/golang-jwt/jwt"

type (
	Token struct {
		jwt.StandardClaims
		Data TokenPayload `json:"data"`
	}

	TokenPayload struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		Timestamp int    `json:"timestamp"`
	}

	repository struct {
		secret string
		issuer string
	}
)
