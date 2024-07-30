package api

import (
	"fmt"
	"net/http"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/models"
)

type FutureClient struct {
	baseUrl    string
	cert       string
	httpClient *http.Client
}

func NewFutureClient(baseUrl, cert string) *FutureClient {
	return &FutureClient{
		baseUrl:    baseUrl,
		cert:       cert,
		httpClient: &http.Client{},
	}
}

func (f *FutureClient) GetInfoByTicker(ticker string) (models.Security, error) {
	res := models.Security{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s.json?iss.only=securities`,
		f.baseUrl,
		ticker,
	)

	resp, err := f.doRequest(url, fmt.Sprintf("failed to get future info by ticker: %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapSecuritySingleResponse()
	return res, nil
}

func (f *FutureClient) GetAllInfo() ([]models.Security, error) {
	res := []models.Security{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities.json?iss.only=securities`,
		f.baseUrl,
	)

	resp, err := f.doRequest(url, "failed to get futures info")
	if err != nil {
		return res, err
	}

	res = resp.MapSecurityFullResponse()
	return res, nil
}

func (f *FutureClient) GetMarketDataByTicker(ticker string) (models.MarketData, error) {
	res := models.MarketData{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s.json?iss.only=marketdata`,
		f.baseUrl,
		ticker,
	)

	resp, err := f.doRequest(url, fmt.Sprintf("failed to get future market data by ticker: %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapMarketDataSingleResponse()
	return res, nil
}

func (f *FutureClient) GetAllMarketData() ([]models.MarketData, error) {
	res := []models.MarketData{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities.json?iss.only=marketdata`,
		f.baseUrl,
	)

	resp, err := f.doRequest(url, "failed to get futures market data")
	if err != nil {
		return res, err
	}

	res = resp.MapMarketDataFullResponse()
	return res, nil
}

func (f *FutureClient) GetCandlesByTicker(ticker, from, till, interval string) ([]models.Candle, error) {
	res := []models.Candle{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s/candles.json?from=%s&till=%s&interval=%s`,
		f.baseUrl,
		ticker,
		from,
		till,
		interval,
	)

	resp, err := f.doRequest(url, fmt.Sprintf("failed to get candles for future: %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapCandleResponse()
	return res, nil
}

func (f *FutureClient) GetOrderBookByTicker(ticker string) ([]models.OrderBook, error) {
	res := []models.OrderBook{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s/orderbook.json`,
		f.baseUrl,
		ticker,
	)

	resp, err := f.doRequest(url, fmt.Sprintf("failed to get orderbook for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapOrderBookResponse()
	return res, nil
}

func (f *FutureClient) GetTradesByTicker(ticker string) ([]models.Trade, error) {
	res := []models.Trade{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s/trades.json`,
		f.baseUrl,
		ticker,
	)

	resp, err := f.doRequest(url, fmt.Sprintf("failed to get trades for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapTradeResponse()
	return res, nil
}
