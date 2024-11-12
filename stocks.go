package moexalgopackgo

import "github.com/DBoyara/MoexAlgoPackGo/pkg/models"

func (a *AlgoClient) GetStockOrderBookByTicker(ticker string) ([]models.OrderBook, error) {
	return a.apiClient.GetStockOrderBookByTicker(ticker)
}
