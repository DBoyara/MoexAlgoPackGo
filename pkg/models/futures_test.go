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

func TestMapTradeData(t *testing.T) {
	m := map[string]interface{}{
		"TRADENO":       1,
		"BOARDNAME":     "RFUD",
		"SECID":         "MMM4",
		"TRADEDATE":     "2022-01-01",
		"TRADETIME":     "10:00:00",
		"PRICE":         100.0,
		"QUANTITY":      10,
		"SYSTIME":       "2022-01-01 10:00:00",
		"RECNO":         257357101010,
		"OPENPOSITION":  76868,
		"OFFMARKETDEAL": 0,
		"BUYSELL":       "B",
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

func TestMapTradeStats(t *testing.T) {
	m := map[string]interface{}{
		"tradedate":    "2022-01-01",
		"tradetime":    "10:00:00",
		"secid":        "MMM4",
		"asset_code":   "MXI",
		"pr_open":      100.0,
		"pr_high":      200.0,
		"pr_low":       50.0,
		"pr_close":     150.0,
		"pr_std":       10.0,
		"vol":          1000,
		"val":          5000,
		"trades":       10,
		"pr_vwap":      150.0,
		"pr_change":    0.5,
		"trades_b":     5,
		"trades_s":     5,
		"val_b":        2500.0,
		"val_s":        2500.0,
		"vol_b":        500,
		"vol_s":        500,
		"disb":         0.0,
		"pr_vwap_b":    150.0,
		"pr_vwap_s":    150.0,
		"im":           10000.0,
		"oi_open":      100,
		"oi_high":      200,
		"oi_low":       50,
		"oi_close":     150,
		"sec_pr_open":  10,
		"sec_pr_high":  20,
		"sec_pr_low":   5,
		"sec_pr_close": 15,
		"SYSTIME":      "2022-01-01 10:00:00",
	}

	expected := TradeStats{
		TradeDate:  "2022-01-01",
		TradeTime:  "10:00:00",
		SecId:      "MMM4",
		AssetCode:  "MXI",
		PrOpen:     100.0,
		PrHigh:     200.0,
		PrLow:      50.0,
		PrClose:    150.0,
		PrStd:      10.0,
		Vol:        1000,
		Val:        5000,
		Trades:     10,
		PrVwap:     150.0,
		PrChange:   0.5,
		TradesB:    5,
		TradesS:    5,
		ValB:       2500.0,
		ValS:       2500.0,
		VolB:       500,
		VolS:       500,
		Disb:       0.0,
		PrVwapB:    150.0,
		PrVwapS:    150.0,
		Im:         10000.0,
		OiOpen:     100,
		OiHigh:     200,
		OiLow:      50,
		OiClose:    150,
		SecPrOpen:  10,
		SecPrHigh:  20,
		SecPrLow:   5,
		SecPrClose: 15,
		SysTime:    "2022-01-01 10:00:00",
	}

	result := mapTradeStats(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapTradeStats() = %v, want %v", result, expected)
	}
}

func TestMapObStats(t *testing.T) {
	m := map[string]interface{}{
		"tradedate":   "2022-01-01",
		"tradetime":   "12:00:00",
		"secid":       "AAPL",
		"asset_code":  "AAPL",
		"mid_price":   100.0,
		"micro_price": 99.5,
		"spread_l1":   0.5,
		"spread_l2":   1.0,
		"spread_l3":   1.5,
		"spread_l5":   2.5,
		"spread_l10":  5.0,
		"spread_l20":  10.0,
		"levels_b":    10,
		"levels_s":    10,
		"vol_b_l1":    100,
		"vol_b_l2":    200,
		"vol_b_l3":    300,
		"vol_b_l5":    500,
		"vol_b_l10":   1000,
		"vol_b_l20":   2000,
		"vol_s_l1":    100,
		"vol_s_l2":    200,
		"vol_s_l3":    300,
		"vol_s_l5":    500,
		"vol_s_l10":   1000,
		"vol_s_l20":   2000,
		"vwap_b_l3":   99.8,
		"vwap_b_l5":   99.7,
		"vwap_b_l10":  99.5,
		"vwap_b_l20":  99.0,
		"vwap_s_l3":   100.2,
		"vwap_s_l5":   100.3,
		"vwap_s_l10":  100.5,
		"vwap_s_l20":  101.0,
		"SYSTIME":     "2022-01-01 12:00:00",
	}

	expected := ObStats{
		TradeDate:  "2022-01-01",
		TradeTime:  "12:00:00",
		SecId:      "AAPL",
		AssetCode:  "AAPL",
		MidPrice:   100.0,
		MicroPrice: 99.5,
		SpreadL1:   0.5,
		SpreadL2:   1.0,
		SpreadL3:   1.5,
		SpreadL5:   2.5,
		SpreadL10:  5.0,
		SpreadL20:  10.0,
		LevelsB:    10,
		LevelsS:    10,
		VolBL1:     100,
		VolBL2:     200,
		VolBL3:     300,
		VolBL5:     500,
		VolBL10:    1000,
		VolBL20:    2000,
		VolSL1:     100,
		VolSL2:     200,
		VolSL3:     300,
		VolSL5:     500,
		VolSL10:    1000,
		VolSL20:    2000,
		VwapBL3:    99.8,
		VwapBL5:    99.7,
		VwapBL10:   99.5,
		VwapBL20:   99.0,
		VwapSL3:    100.2,
		VwapSL5:    100.3,
		VwapSL10:   100.5,
		VwapSL20:   101.0,
		SysTime:    "2022-01-01 12:00:00",
	}

	result := mapObStats(m)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("mapObStats returned unexpected result, got: %v, want: %v", result, expected)
	}
}
