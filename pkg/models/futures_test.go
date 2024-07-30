package models

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"

	"github.com/DBoyara/MoexAlgoPackGo/pkg/utils"
)

var GetInfoByTicker = Response{
	Securities: Data{
		Columns: []string{
			"SECID",
			"BOARDID",
			"SHORTNAME",
			"SECNAME",
			"PREVSETTLEPRICE",
			"DECIMALS",
			"MINSTEP",
			"LASTTRADEDATE",
			"LASTDELDATE",
			"SECTYPE",
			"LATNAME",
			"ASSETCODE",
			"PREVOPENPOSITION",
			"LOTVOLUME",
			"INITIALMARGIN",
			"HIGHLIMIT",
			"LOWLIMIT",
			"STEPPRICE",
			"LASTSETTLPRICE",
			"PREVPRICE",
			"IMTIME",
			"BUYSELLFEE",
			"SCALPERFEE",
			"NEGOTIATEDFEE",
			"EXERCISEFEE",
		},
		Data: [][]interface{}{
			{
				"MMM4",
				"RFUD",
				"MXI-6.24",
				"Фьючерсный контракт MXI-6.24",
				3462.65,
				2.0,
				0.05000,
				"2024-06-20",
				"2024-06-20",
				"MM",
				"MXI-6.24",
				"MXI",
				76868.0,
				1.0,
				5804.59,
				3748.70000,
				3176.60000,
				0.50000,
				3462.2,
				3462.2,
				"2024-04-11 18:58:42",
				2.28,
				1.14,
				0.76,
				0.76,
			},
		},
	},
}

func BenchmarkMapResponse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetInfoByTicker.MapSecuritySingleResponse()
	}
}

func BenchmarkMapingResponse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetInfoByTicker.MapingResponseByJson()
	}
}

func (r *Response) MapingResponseByJson() *Security {
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
	m := data[0]
	s := &Security{}

	jsonData, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error marshaling map: %v", err)
	}

	if err := json.Unmarshal(jsonData, &s); err != nil {
		log.Fatalf("Error unmarshaling JSON into struct: %v", err)
	}

	return s
}

func TestMapSecirity(t *testing.T) {
	m := map[string]interface{}{
		"SECID":            "MMM4",
		"BOARDID":          "RFUD",
		"SHORTNAME":        "MXI-6.24",
		"SECNAME":          "Фьючерсный контракт MXI-6.24",
		"PREVSETTLEPRICE":  3462.65,
		"DECIMALS":         2,
		"MINSTEP":          0.05000,
		"LASTTRADEDATE":    "2024-06-20",
		"LASTDELDATE":      "2024-06-20",
		"SECTYPE":          "MM",
		"LATNAME":          "MXI-6.24",
		"ASSETCODE":        "MXI",
		"PREVOPENPOSITION": 76868,
		"LOTVOLUME":        1,
		"INITIALMARGIN":    5804.59,
		"HIGHLIMIT":        3748.70000,
		"LOWLIMIT":         3176.60000,
		"STEPPRICE":        0.50000,
		"LASTSETTLPRICE":   3462.2,
		"PREVPRICE":        3462.2,
		"IMTIME":           "2024-04-11 18:58:42",
		"BUYSELLFEE":       2.28,
		"SCALPERFEE":       1.14,
		"NEGOTIATEDFEE":    0.76,
		"EXERCISEFEE":      0.76,
	}

	expected := Security{
		SecId:               "MMM4",
		BoardId:             "RFUD",
		ShortName:           "MXI-6.24",
		SecName:             "Фьючерсный контракт MXI-6.24",
		PrevSettlePrice:     3462.65,
		Decimals:            2,
		MinStep:             0.05000,
		LastTradeDate:       "2024-06-20",
		LastDelDate:         "2024-06-20",
		SecType:             "MM",
		LatName:             "MXI-6.24",
		AssetCode:           "MXI",
		PreviosOpenPosition: 76868,
		LotVolume:           1,
		InitalMargin:        5804.59,
		HighLimit:           3748.70000,
		LowLimit:            3176.60000,
		StepPrice:           0.50000,
		LastSettlPrice:      3462.2,
		PrevPrice:           3462.2,
		ImTime:              "2024-04-11 18:58:42",
		BuySellFee:          2.28,
		ScalperFee:          1.14,
		NegotiatedFee:       0.76,
		ExerciseFee:         0.76,
	}

	result := mapSecirity(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapData() = %v, want %v", result, expected)
	}
}

func TestMapMarketData(t *testing.T) {
	m := map[string]interface{}{
		"SECID":                 "MMM4",
		"BOARDID":               "RFUD",
		"BID":                   3462.65,
		"OFFER":                 3462.7,
		"SPREAD":                0.05,
		"OPEN":                  3462.2,
		"HIGH":                  3462.2,
		"LOW":                   3462.2,
		"LAST":                  3462.2,
		"QUANTITY":              1,
		"LASTCHANGE":            0.3,
		"SETTLEPRICE":           3462.2,
		"SETTLETOPREVSETTLE":    2.25,
		"OPENPOSITION":          0,
		"NUMTRADES":             5159,
		"VOLTODAY":              15966,
		"VALTODAY":              553262996,
		"VALTODAY_USD":          5935037,
		"UPDATETIME":            "18:58:42",
		"LASTCHANGEPRCNT":       0.01,
		"BIDDEPTH":              1,
		"BIDDEPTHT":             5329,
		"NUMBIDS":               709,
		"OFFERDEPTH":            14,
		"OFFERDEPTHT":           2879,
		"NUMOFFERS":             425,
		"TIME":                  "18:58:42",
		"SETTLETOPREVSETTLEPRC": 0,
		"SEQNUM":                1,
		"SYSTIME":               "2024-04-11 18:58:42",
		"TRADEDATE":             "2024-04-11 18:58:42",
		"LASTTOPREVPRICE":       0.1,
		"OICHANGE":              -638,
		"OPENPERIODPRICE":       3462.75000,
		"SWAPRATE":              0.0,
	}

	expected := MarketData{
		SecId:                 "MMM4",
		BoardId:               "RFUD",
		Bid:                   3462.65,
		Offer:                 3462.7,
		Spread:                0.05,
		Open:                  3462.2,
		High:                  3462.2,
		Low:                   3462.2,
		Last:                  3462.2,
		Quantity:              1,
		LastChange:            0.3,
		SettlePrice:           3462.2,
		SettleToPrevSettle:    2.25,
		OpenPosition:          0,
		NumTrades:             5159,
		VolToday:              15966,
		ValToday:              553262996,
		ValTodayUsd:           5935037,
		UpdateTime:            "18:58:42",
		LastChangePrcnt:       0.01,
		BidDepth:              1,
		BidDepthT:             5329,
		NumBids:               709,
		OfferDepth:            14,
		OfferDepthT:           2879,
		NumOffers:             425,
		Time:                  "18:58:42",
		SettleToPrevSettlePrc: 0,
		SeqNum:                1,
		SysTime:               "2024-04-11 18:58:42",
		TradeDate:             "2024-04-11 18:58:42",
		LastToPrevPrice:       0.1,
		OiChange:              -638,
		OpenPeriodPrice:       3462.75000,
		SwapRate:              0.0,
	}

	result := mapMarketData(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapMarketData() = %v, want %v", result, expected)
	}
}

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
		"BOARDID":   "RFUD",
		"SECID":     "MMM4",
		"BUYSELL":   "B",
		"PRICE":     3462.65,
		"QUANTITY":  100,
		"SEQNUM":    20240412110944,
		"UPDATETIME": "18:58:42",
		"DECIMALS":  2,
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

func TestMapTradeData(t *testing.T) {
	m := map[string]interface{}{
		"TRADENO":        1,
		"BOARDNAME":      "RFUD",
		"SECID":          "MMM4",
		"TRADEDATE":      "2022-01-01",
		"TRADETIME":      "10:00:00",
		"PRICE":          100.0,
		"QUANTITY":       10,
		"SYSTIME":        "2022-01-01 10:00:00",
		"RECNO":          257357101010,
		"OPENPOSITION":   76868,
		"OFFMARKETDEAL":  0,
		"BUYSELL":        "B",
	}

	expected := Trade{
		TradeNo:       1,
		BoardName:     "RFUD",
		SecId:         "MMM4",
		TradeDate:     "2022-01-01",
		TradeTime:     "10:00:00",
		Price:         100.0,
		Quantity:      10,
		SysTime:       "2022-01-01 10:00:00",
		RecNo:         257357101010,
		OpenPosition:  76868,
		OffMarketDeal: 0,
		BuySell:       "B",
	}

	result := mapTradeData(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapTradeData() = %v, want %v", result, expected)
	}
}