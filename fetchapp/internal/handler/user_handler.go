package handler

import (
	"net/http"

	"github.com/marprin/assessment/fetchapp/pkg/response"
)

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	resp := h.userService.Profile(r.Context())

	response.Response(w, r, http.StatusOK, resp, nil, nil, http.StatusOK)
}
