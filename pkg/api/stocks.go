package api

import (
	"fmt"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/models"
)

func (a *ApiClient) GetStockOrderBookByTicker(ticker string) ([]models.OrderBook, error) {
	res := []models.OrderBook{}
	url := fmt.Sprintf(
		`%s/iss/engines/stock/markets/shares/boards/tqbr/securities/%s/orderbook.json`,
		a.baseUrl,
		ticker,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to get orderbook for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapOrderBookResponse()
	return res, nil
}
