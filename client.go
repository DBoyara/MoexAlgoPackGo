package moexalgopackgo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/api"
	"github.com/DBoyara/MoexAlgoPackGo/pkg/models"
)

type IAlgoClient interface {
	// Торговые параметры о фьючерсе (наименорвание, шаг цены, лотность, ...)
	GetFutureInfoByTicker(ticker string) (models.Security, error)
	// Торговые параметры о всех фьючерсах (наименорвание, шаг цены, лотность, ...)
	GetFuturesInfo() ([]models.Security, error)
	// Торговая статистика по одному фьючерсу за текущий день (цены, объемы, ...)
	GetFutureMarketDataByTicker(ticker string) (models.MarketData, error)
	// Торговая статистика по всем фьючерсам за текущий день (цены, объемы, ...)
	GetFuturesMarketData() ([]models.MarketData, error)
	// Свечи по одному указанному фьючерсному контракту
	GetFuturesCandles(ticker, from, till, interval string) ([]models.Candle, error)
	// Стакан котировок по одному указанному фьючерсному контракту
	GetFuturesOrderBookByTiker(ticker string) ([]models.OrderBook, error)
	// Метрики рассчитанные на основе потока сделок (tradestats) по одному фьючерсному контракту
	GetFuturesTradeStatsByTiker(ticker, from, till, latest string) ([]models.TradeStats, error)
	// Метрики рассчитанные на основе стакана котировок (obstats) по одному фьючерсному контракту
	GetFuturesObStatsByTiker(ticker, from, till, latest string) ([]models.ObStats, error)
}

type AlgoClient struct {
	auth_url string
	base_url string
	login    string
	password string
	cert     string
	futures  *api.FutureClient
}

func NewAlgoClient(login, password string) IAlgoClient {
	client := &AlgoClient{
		auth_url: "https://passport.moex.com/authenticate",
		base_url: "https://iss.moex.com",
		login:    login,
		password: password,
	}

	err := client.authenticate()
	if err != nil {
		log.Fatalf("failed to authenticate: %v", err)
	}

	client.futures = api.NewFutureClient(client.base_url, client.cert)

	return client
}

func (a *AlgoClient) authenticate() error {
	request, err := http.NewRequest(http.MethodGet, a.auth_url, nil)
	if err != nil {
		return err
	}

	request.SetBasicAuth(a.login, a.password)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to authenticate | Status Code: %d | Text: %s", response.StatusCode, response.Status)
	}

	cookies := response.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "MicexPassportCert" {
			a.cert = cookie.Value
			break
		}
	}

	return nil
}
