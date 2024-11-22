package api

import (
	"fmt"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/models"
)

func (a *ApiClient) GetInfoByTicker(ticker string) (models.Security, error) {
	res := models.Security{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s.json?iss.only=securities`,
		a.baseUrl,
		ticker,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to get future info by ticker: %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapSecuritySingleResponse()
	return res, nil
}

func (a *ApiClient) GetAllInfo() ([]models.Security, error) {
	res := []models.Security{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities.json?iss.only=securities`,
		a.baseUrl,
	)

	resp, err := a.doRequest(url, "failed to get futures info")
	if err != nil {
		return res, err
	}

	res = resp.MapSecurityFullResponse()
	return res, nil
}

func (a *ApiClient) GetMarketDataByTicker(ticker string) (models.MarketData, error) {
	res := models.MarketData{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s.json?iss.only=marketdata`,
		a.baseUrl,
		ticker,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to get future market data by ticker: %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapMarketDataSingleResponse()
	return res, nil
}

func (a *ApiClient) GetAllMarketData() ([]models.MarketData, error) {
	res := []models.MarketData{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities.json?iss.only=marketdata`,
		a.baseUrl,
	)

	resp, err := a.doRequest(url, "failed to get futures market data")
	if err != nil {
		return res, err
	}

	res = resp.MapMarketDataFullResponse()
	return res, nil
}

func (a *ApiClient) GetTradesByTicker(ticker string) ([]models.Trade, error) {
	res := []models.Trade{}
	url := fmt.Sprintf(
		`%s/iss/engines/futures/markets/forts/boards/rfud/securities/%s/trades.json`,
		a.baseUrl,
		ticker,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to get trades for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.MapTradeResponse()
	return res, nil
}

func (a *ApiClient) GetFuturesTradeStatsByTicker(ticker, from, till, latest string) ([]models.TradeStats, error) {
	res := []models.TradeStats{}
	url := fmt.Sprintf(
		`%s/iss/datashop/algopack/fo/tradestats/%s.json?from=%s&till=%s&latest=%s`,
		a.baseUrl,
		ticker,
		from,
		till,
		latest,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to tradestats for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.TradeStatsResponse()
	return res, nil
}

func (a *ApiClient) GetFuturesObStatsByTicker(ticker, from, till, latest string) ([]models.ObStats, error) {
	res := []models.ObStats{}
	url := fmt.Sprintf(
		`%s/iss/datashop/algopack/fo/obstats/%s.json?from=%s&till=%s&latest=%s`,
		a.baseUrl,
		ticker,
		from,
		till,
		latest,
	)

	resp, err := a.doRequest(url, fmt.Sprintf("failed to obstats for future %s", ticker))
	if err != nil {
		return res, err
	}

	res = resp.ObStatsResponse()
	return res, nil
}
