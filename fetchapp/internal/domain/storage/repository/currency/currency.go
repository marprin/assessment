package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/marprin/assessment/fetchapp/internal/domain/storage"
	"github.com/marprin/assessment/fetchapp/internal/domain/storage/entity"
	"github.com/sirupsen/logrus"
)

type (
	currency struct {
		url    string
		apiKey string
	}
)

func New(url, apiKey string) storage.Currency {
	return &currency{
		url:    url,
		apiKey: apiKey,
	}
}

func (c *currency) FetchCurrencyConverter(ctx context.Context, q string) (*float64, error) {
	url := fmt.Sprintf("%s?apiKey=%s&q=%s&compact=y", c.url, c.apiKey, q)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchCurrencyConverter] Error when make http request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchCurrencyConverter] Error when do http request")
		return nil, err
	}

	defer resp.Body.Close()

	returnResp := entity.CurrencyConverterResponse{}
	err = json.NewDecoder(resp.Body).Decode(&returnResp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorln("[FetchCurrencyConverter] Error when decode response")
		return nil, err
	}

	return &returnResp.USDIDR.Val, nil
}
