package models

import (
	"reflect"
	"testing"
)

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
