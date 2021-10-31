package handler

import (
	"github.com/marprin/assessment/fetchapp/internal/domain/storage"
	"github.com/marprin/assessment/fetchapp/internal/domain/user"
)

type (
	Handler struct {
		userService    user.Service
		storageService storage.Service
	}

	Options struct {
		UserService    user.Service
		StorageService storage.Service
	}
)

func New(o *Options) *Handler {
	return &Handler{
		userService:    o.UserService,
		storageService: o.StorageService,
	}
}
