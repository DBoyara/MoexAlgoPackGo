package moexalgopackgo

import "github.com/DBoyara/MoexAlgoPackGo/pkg/models"

const stockBoard = "tqbr"

func (a *AlgoClient) GetStockOrderBookByTicker(ticker string) ([]models.OrderBook, error) {
	return a.apiClient.GetOrderBookByTicker(ticker, stockBoard)
}

func (a *AlgoClient) GetStockCandlesByTicker(ticker, from, till, interval string) ([]models.Candle, error) {
	return a.apiClient.GetCandlesByTicker(ticker, from, till, interval, stockBoard)
}
