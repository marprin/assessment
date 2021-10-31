package handler

import (
	"net/http"

	"github.com/marprin/assessment/fetchapp/internal/constant"
	"github.com/marprin/assessment/fetchapp/pkg/response"
)

func (h *Handler) FetchStorage(w http.ResponseWriter, r *http.Request) {

	resp, err := h.storageService.StorageList(r.Context())
	if err != nil {
		errRes := response.ErrorPayload{
			Message: constant.ErrInternalServer.Error(),
		}
		response.Response(w, r, http.StatusInternalServerError, nil, errRes, nil, http.StatusInternalServerError)
		return
	}

	response.Response(w, r, http.StatusOK, resp, nil, nil, http.StatusOK)
	return
}
