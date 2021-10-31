package entity

import "time"

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
)
