package jwt

import (
	"context"
	"time"

	jwtLib "github.com/dgrijalva/jwt-go"
	"github.com/marprin/assessment/fetchapp/internal/constant"
)

func New(secret, issuer string) Repository {
	return &repository{
		secret: secret,
		issuer: issuer,
	}
}

func (r *repository) ExtractToken(ctx context.Context, tokenJwt string) (*TokenPayload, error) {
	token, err := jwtLib.ParseWithClaims(tokenJwt, &Token{}, func(token *jwtLib.Token) (interface{}, error) {
		return []byte(r.secret), nil
	})
	if err != nil && err.Error() == jwtLib.ErrInvalidKeyType.Error() {
		return nil, constant.ErrTokenIsNotValid
	}

	claims, ok := token.Claims.(*Token)
	if !ok || claims == nil {
		return nil, constant.ErrClaimsIsNotValid
	}

	if claims.Issuer != r.issuer {
		return nil, constant.ErrIssuerIsNotValid
	}

	if claims.ExpiresAt <= time.Now().Unix() {
		return nil, constant.ErrTokenIsExpired
	}

	return &TokenPayload{
		Name:      claims.Data.Name,
		Phone:     claims.Data.Phone,
		Role:      claims.Data.Role,
		Timestamp: claims.Data.Timestamp,
	}, nil

}
