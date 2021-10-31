package user

import (
	"context"

	"github.com/marprin/assessment/fetchapp/internal/domain/user/entity"
)

type (
	Service interface {
		Profile(ctx context.Context) *entity.ProfileResponse
	}
)
