package entity

import (
	"time"

	"github.com/marprin/assessment/fetchapp/internal/constant"
)

type (
	StorageListResponse struct {
		UUID         *string    `json:"uuid"`
		Komoditas    *string    `json:"komoditas"`
		AreaProvinsi *string    `json:"area_provinsi"`
		AreaKota     *string    `json:"area_kota"`
		Size         *string    `json:"size"`
		Price        *string    `json:"price"`
		TglParsed    *time.Time `json:"tgl_parsed"`
		Timestamp    *string    `json:"timestamp"`
		USDPrice     *float64   `json:"usd_price"`
	}

	FilterStorageRequest struct {
		AreaProvinsi  string
		StartDate     string
		EndDate       string
		StartDateTime int
		EndDateTime   int
	}
	FilterStorageResponse struct {
		Min    float64 `json:"min"`
		Max    float64 `json:"max"`
		Median float64 `json:"median"`
		Avg    float64 `json:"avg"`
	}
)

func (f *FilterStorageRequest) Validate() error {
	if f.StartDate != "" {
		parseStartTime, err := time.Parse("2006-01-02T15:04:05", f.StartDate)
		if err != nil {
			return constant.ErrStartDateNotValid
		}
		f.StartDateTime = int(parseStartTime.Unix())
	}

	if f.EndDate != "" {
		parseEndTime, err := time.Parse("2006-01-02T15:04:05", f.EndDate)
		if err != nil {
			return constant.ErrEndDateNotValid
		}
		f.EndDateTime = int(parseEndTime.Unix())
	}

	return nil
}
