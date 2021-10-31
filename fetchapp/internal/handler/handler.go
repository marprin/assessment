package handler

import "github.com/marprin/assessment/fetchapp/internal/domain/user"

type (
	Handler struct {
		userService user.Service
	}

	Options struct {
		UserService user.Service
	}
)

func New(o *Options) *Handler {
	return &Handler{
		userService: o.UserService,
	}
}
