package service

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/marprin/assessment/fetchapp/internal/domain/storage"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/entity"
	"github.com/marprin/assessment/fetchapp/pkg/cache"
	"github.com/sirupsen/logrus"
)

type (
	service struct {
		currencyRepo storage.Currency
		gatewayRepo  storage.Gateway
		cache        cache.Cache
	}

	Options struct {
		CurrencyRepo storage.Currency
		GatewayRepo  storage.Gateway
		Cache        cache.Cache
	}
)

func New(o *Options) storage.Service {
	return &service{
		currencyRepo: o.CurrencyRepo,
		gatewayRepo:  o.GatewayRepo,
		cache:        o.Cache,
	}
}

func (s *service) StorageList(ctx context.Context) ([]entity.StorageListResponse, error) {

	gatewayResp, err := s.gatewayRepo.FetchStorageList(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[StorageList, FetchStorageList] Error when fetch storage list")
		return nil, err
	}

	var usdIDR float64
	shouldCheck := true
	cachedValue, err := s.cache.Get("./data/currency.txt")
	if err == nil && len(cachedValue) > 0 {
		usdIDR, err = strconv.ParseFloat(cachedValue[0], 64)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("[StorageList, ParseFloat] Error when parse to float")
			return nil, err
		}
		shouldCheck = false
	}

	if shouldCheck {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[StorageList, Get] Error when fetch cached currency")

		usdIDRResp, err := s.currencyRepo.FetchCurrencyConverter(ctx, "USD_IDR")
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("[StorageList, FetchCurrencyConverter] Error when fetch currency exchange value")
			return nil, err
		}
		usdIDR = *usdIDRResp

		err = s.cache.Set("./data/currency.txt", fmt.Sprintf("%f", usdIDR))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"err": err,
			}).Errorln("[StorageList, Set] Error when set to cache for currency exchange value")
			return nil, err
		}
	}

	resp := []entity.StorageListResponse{}
	for _, v := range gatewayResp {
		if v.UUID == nil {
			continue
		}

		var usdPrice *float64
		if v.Price != nil {
			priceFloat, err := strconv.ParseFloat(*v.Price, 64)
			if err == nil {
				price := priceFloat / usdIDR
				usdPrice = &price
			}
		}

		resp = append(resp, entity.StorageListResponse{
			UUID:         v.UUID,
			Komoditas:    v.Komoditas,
			AreaProvinsi: v.AreaProvinsi,
			AreaKota:     v.AreaKota,
			Size:         v.Size,
			Price:        v.Price,
			TglParsed:    v.TglParsed,
			Timestamp:    v.Timestamp,
			USDPrice:     usdPrice,
		})
	}

	return resp, nil
}

func (s *service) FilterStorageList(ctx context.Context, payload entity.FilterStorageRequest) (*entity.FilterStorageResponse, error) {

	gatewayResp, err := s.gatewayRepo.FetchStorageList(ctx)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[StorageList, FilterStorageList] Error when fetch storage list")
		return nil, err
	}

	var (
		minPrice, maxPrice, sumPrice float64 = math.MaxFloat64, -1, 0
		totalItem                    int
	)
	for _, v := range gatewayResp {
		if v.UUID == nil || v.Price == nil {
			continue
		}

		if payload.AreaProvinsi != "" && *v.AreaProvinsi != payload.AreaProvinsi {
			continue
		}

		if v.Timestamp != nil {
			parsedTime, err := strconv.Atoi(*v.Timestamp)
			if err != nil {
				continue
			}

			if payload.StartDate != "" && parsedTime < payload.StartDateTime {
				continue
			}

			if payload.EndDate != "" && parsedTime > payload.EndDateTime {
				continue
			}
		}

		priceFloat, err := strconv.ParseFloat(*v.Price, 64)
		if err != nil {
			continue
		}

		if priceFloat > maxPrice {
			maxPrice = priceFloat
		}
		if priceFloat < minPrice {
			minPrice = priceFloat
		}

		sumPrice += priceFloat

		totalItem++
	}

	median := float64(0)
	avg := float64(0)
	if sumPrice != 0 {
		median = sumPrice / 2

		if totalItem > 0 {
			avg = sumPrice / float64(totalItem)
		}
	}

	return &entity.FilterStorageResponse{
		Min:    minPrice,
		Max:    maxPrice,
		Median: median,
		Avg:    avg,
	}, nil

}
