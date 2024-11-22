package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestFutureClient_GetCandlesByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"candles": {"columns": ["open", "close", "high", "low", "value", "volume", "begin", "end"], "data": [[100849, 100950, 100950, 100849, 0, 12, "2024-04-15 10:00:00", "2024-04-15 10:47:02"]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	candles, err := client.GetCandlesByTicker("sber", "2022-01-01", "2022-01-02", "1", "tqbr")

	assert.NoError(t, err)
	assert.Len(t, candles, 1)
}

func TestFutureClient_GetOrderBookByTicker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"orderbook": {"columns": ["BOARDID", "SECID", "BUYSELL", "PRICE", "QUANTITY", "SEQNUM", "UPDATETIME", "DECIMALS"], "data": [["RFUD", "SiH5", "B", 100097, 1, 20240416195149, "19:51:42", 0]]}}`))
	}))
	defer server.Close()

	client := api.NewApiClient(server.URL, "")
	orderBook, err := client.GetOrderBookByTicker("RFUD", "tqbr")

	assert.NoError(t, err)
	assert.Len(t, orderBook, 1)
}
