package handler

import (
	"net/http"

	"github.com/marprin/assessment/fetchapp/internal/constant"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/entity"
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

func (h *Handler) FilterStorage(w http.ResponseWriter, r *http.Request) {
	payload := entity.FilterStorageRequest{
		AreaProvinsi: r.URL.Query().Get("area_provinsi"),
		StartDate:    r.URL.Query().Get("start_date"),
		EndDate:      r.URL.Query().Get("end_date"),
	}

	err := payload.Validate()
	if err != nil {
		errRes := response.ErrorPayload{
			Message: err.Error(),
		}
		response.Response(w, r, http.StatusBadRequest, nil, errRes, nil, http.StatusBadRequest)
		return
	}

	resp, err := h.storageService.FilterStorageList(r.Context(), payload)
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
