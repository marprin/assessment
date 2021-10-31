package service

import (
	"context"

	"github.com/marprin/assessment/fetchapp/internal/config"
	"github.com/marprin/assessment/fetchapp/internal/domain/user"
	"github.com/marprin/assessment/fetchapp/internal/domain/user/entity"
	"github.com/marprin/assessment/fetchapp/internal/middleware"
	"github.com/marprin/assessment/fetchapp/pkg/jwt"
)

type (
	service struct {
		config *config.Config
		jwt    jwt.Repository
	}
)

func New(cfg *config.Config, jwt jwt.Repository) user.Service {
	return &service{
		config: cfg,
		jwt:    jwt,
	}
}

func (s *service) Profile(ctx context.Context) *entity.ProfileResponse {

	return &entity.ProfileResponse{
		Name:  ctx.Value(middleware.NameCtx).(string),
		Phone: ctx.Value(middleware.PhoneCtx).(string),
		Role:  ctx.Value(middleware.RoleCtx).(string),
	}
}
