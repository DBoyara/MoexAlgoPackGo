package moexalgopackgo

import "github.com/DBoyara/MoexAlgoPackGo/pkg/models"

func (a *AlgoClient) GetFutureInfoByTicker(ticker string) (models.Security, error) {
	return a.futures.GetInfoByTicker(ticker)
}

func (a *AlgoClient) GetFuturesInfo() ([]models.Security, error) {
	return a.futures.GetAllInfo()
}

func (a *AlgoClient) GetFutureMarketDataByTicker(ticker string) (models.MarketData, error) {
	return a.futures.GetMarketDataByTicker(ticker)
}

func (a *AlgoClient) GetFuturesMarketData() ([]models.MarketData, error) {
	return a.futures.GetAllMarketData()
}

func (a *AlgoClient) GetFuturesCandles(ticker, from, till, interval string) ([]models.Candle, error) {
	return a.futures.GetCandlesByTicker(ticker, from, till, interval)
}

func (a *AlgoClient) GetFuturesOrderBookByTiker(ticker string) ([]models.OrderBook, error) {
	return a.futures.GetOrderBookByTicker(ticker)
}
