package models

import (
	"reflect"
	"testing"
)

func TestMapCandleData(t *testing.T) {
	m := map[string]interface{}{
		"open":   100.0,
		"close":  150.0,
		"high":   200.0,
		"low":    50.0,
		"value":  1000,
		"volume": 500.0,
		"begin":  "2022-01-26 23:20:00",
		"end":    "2022-01-26 23:27:01",
	}

	expected := Candle{
		Open:   100.0,
		Close:  150.0,
		High:   200.0,
		Low:    50.0,
		Value:  1000,
		Volume: 500.0,
		Begin:  "2022-01-26 23:20:00",
		End:    "2022-01-26 23:27:01",
	}

	result := mapCandleData(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapCandleData() = %v, want %v", result, expected)
	}
}

func TestMapOrderBookData(t *testing.T) {
	m := map[string]interface{}{
		"BOARDID":    "RFUD",
		"SECID":      "MMM4",
		"BUYSELL":    "B",
		"PRICE":      3462.65,
		"QUANTITY":   100,
		"SEQNUM":     20240412110944,
		"UPDATETIME": "18:58:42",
		"DECIMALS":   2,
	}

	expected := OrderBook{
		BoardId:    "RFUD",
		SecId:      "MMM4",
		BuySell:    "B",
		Price:      3462.65,
		Quantity:   100,
		SeqNum:     20240412110944,
		UpdateTime: "18:58:42",
		Decimals:   2,
	}

	result := mapOrderBookData(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapOrderBookData() = %v, want %v", result, expected)
	}
}
