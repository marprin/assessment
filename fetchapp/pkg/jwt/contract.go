package jwt

import "context"

type (
	Repository interface {
		ExtractToken(ctx context.Context, tokenJwt string) (*TokenPayload, error)
	}
)
