package moexalgopackgo

import "github.com/DBoyara/MoexAlgoPackGo/pkg/models"

const futBoard = "fut"

func (a *AlgoClient) GetFutureInfoByTicker(ticker string) (models.Security, error) {
	return a.apiClient.GetInfoByTicker(ticker)
}

func (a *AlgoClient) GetFuturesInfo() ([]models.Security, error) {
	return a.apiClient.GetAllInfo()
}

func (a *AlgoClient) GetFutureMarketDataByTicker(ticker string) (models.MarketData, error) {
	return a.apiClient.GetMarketDataByTicker(ticker)
}

func (a *AlgoClient) GetFuturesMarketData() ([]models.MarketData, error) {
	return a.apiClient.GetAllMarketData()
}

func (a *AlgoClient) GetFuturesCandles(ticker, from, till, interval string) ([]models.Candle, error) {
	return a.apiClient.GetCandlesByTicker(ticker, from, till, interval, futBoard)
}

func (a *AlgoClient) GetFuturesOrderBookByTicker(ticker string) ([]models.OrderBook, error) {
	return a.apiClient.GetOrderBookByTicker(ticker, futBoard)
}

func (a *AlgoClient) GetFuturesTradeStatsByTicker(ticker, from, till, latest string) ([]models.TradeStats, error) {
	return a.apiClient.GetFuturesTradeStatsByTicker(ticker, from, till, latest)
}

func (a *AlgoClient) GetFuturesObStatsByTicker(ticker, from, till, latest string) ([]models.ObStats, error) {
	return a.apiClient.GetFuturesObStatsByTicker(ticker, from, till, latest)
}
