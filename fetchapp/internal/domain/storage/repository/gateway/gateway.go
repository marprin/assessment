package gateway

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/marprin/assessment/fetchapp/internal/domain/storage"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/entity"
	"github.com/sirupsen/logrus"
)

type (
	gateway struct {
		url string
	}
)

func New(url string) storage.Gateway {
	return &gateway{
		url: url,
	}
}

func (g *gateway) FetchStorageList(ctx context.Context) ([]entity.GatewayListResponse, error) {
	req, err := http.NewRequest("GET", g.url, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchStorageList] Error when make http request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchStorageList] Error when do http request")
		return nil, err
	}

	defer resp.Body.Close()

	returnResp := []entity.GatewayListResponse{}
	err = json.NewDecoder(resp.Body).Decode(&returnResp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchStorageList] Error when decode response")
		return nil, err
	}
	return returnResp, nil
}
