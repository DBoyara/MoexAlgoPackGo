package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestFutureClient_GetInfoByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"securities": {"columns": ["SECID"], "data": [["sber"]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	security, err := client.GetInfoByTicker("sber")

	assert.NoError(t, err)
	assert.Equal(t, "sber", security.SecId)
}

func TestFutureClient_GetMarketDataByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"marketdata": {"columns": ["SECID"], "data": [["sber"]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	marketData, err := client.GetMarketDataByTicker("sber")

	assert.NoError(t, err)
	assert.Equal(t, "sber", marketData.SecId)
}

func TestFutureClient_GetTradesByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"trades": {"columns": ["TRADENO"], "data": [[1892949429179121665]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	trades, err := client.GetTradesByTicker("test")

	assert.NoError(t, err)
	assert.Len(t, trades, 1)
}

func TestFutureClient_GetFuturesTradeStatsByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": {"columns": ["SECID"], "data": [["sber"]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	tradeStats, err := client.GetFuturesTradeStatsByTicker("test", "2022-01-01", "2022-01-02", "latest")

	assert.NoError(t, err)
	assert.Len(t, tradeStats, 1)
}

func TestFutureClient_GetFuturesObStatsByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": {"columns": ["SECID"], "data": [["sber"]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	obStats, err := client.GetFuturesObStatsByTicker("test", "2022-01-01", "2022-01-02", "latest")

	assert.NoError(t, err)
	assert.Len(t, obStats, 1)
}
