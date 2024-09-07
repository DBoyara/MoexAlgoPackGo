package models

import (
	"github.com/DBoyara/MoexAlgoPackGo/pkg/utils"
)

type Response struct {
	Securities Data `json:"securities,omitempty"`
	MarketData Data `json:"marketdata,omitempty"`
	Candles    Data `json:"candles,omitempty"`
	OrderBooks Data `json:"orderbook,omitempty"`
	Trades     Data `json:"trades,omitempty"`
	Data       Data `json:"data,omitempty"` // super candles
}

type Data struct {
	Columns []string        `json:"columns"`
	Data    [][]interface{} `json:"data"`
}

type Security struct {
	SecId               string  `json:"sec_id"`                 // Идентификатор ценной бумаги.
	BoardId             string  `json:"board_id"`               // Идентификатор рынка.
	ShortName           string  `json:"short_name"`             // Краткое название ценной бумаги.
	SecName             string  `json:"sec_name"`               // Полное название ценной бумаги.
	PrevSettlePrice     float64 `json:"prev_settle_price"`      // Предыдущая цена закрытия.
	Decimals            int     `json:"decimals"`               // Количество десятичных знаков.
	MinStep             float64 `json:"min_step"`               // Минимальный шаг изменения цены.
	LastTradeDate       string  `json:"last_trade_date"`        // Дата последней сделки.
	LastDelDate         string  `json:"last_del_date"`          // Дата последнего исполнения.
	SecType             string  `json:"sec_type"`               // Тип ценной бумаги.
	LatName             string  `json:"lat_name"`               // Латинское название ценной бумаги.
	AssetCode           string  `json:"asset_code"`             // Код актива.
	PreviosOpenPosition int     `json:"previous_open_position"` // Предыдущая открытая позиция.
	LotVolume           int     `json:"lot_volume"`             // Объем лота.
	InitalMargin        float64 `json:"initial_margin"`         // Начальная маржа.
	HighLimit           float64 `json:"high_limit"`             // Верхний предел цены.
	LowLimit            float64 `json:"low_limit"`              // Нижний предел цены.
	StepPrice           float64 `json:"step_price"`             // Цена шага.
	LastSettlPrice      float64 `json:"last_settl_price"`       // Последняя цена закрытия.
	PrevPrice           float64 `json:"prev_price"`             // Предыдущая цена.
	ImTime              string  `json:"im_time"`                // Время обновления начальной маржи.
	BuySellFee          float64 `json:"buy_sell_fee"`           // Комиссия за покупку/продажу.
	ScalperFee          float64 `json:"scalper_fee"`            // Комиссия за скальпирование.
	NegotiatedFee       float64 `json:"negotiated_fee"`         // Комиссия за договоренность.
	ExerciseFee         float64 `json:"exercise_fee"`           // Комиссия за исполнение.
}

func (r *Response) MapSecuritySingleResponse() Security {
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
	m := data[0]
	return mapSecirity(m)
}

func (r *Response) MapSecurityFullResponse() []Security {
	res := []Security{}
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
	for _, m := range data {
		res = append(res, mapSecirity(m))
	}
	return res
}

func mapSecirity(m map[string]interface{}) Security {
	s := Security{}

	if val, ok := m["SECID"].(string); ok {
		s.SecId = val
	}
	if val, ok := m["BOARDID"].(string); ok {
		s.BoardId = val
	}
	if val, ok := m["SHORTNAME"].(string); ok {
		s.ShortName = val
	}
	if val, ok := m["SECNAME"].(string); ok {
		s.SecName = val
	}
	if val, ok := m["PREVSETTLEPRICE"].(float64); ok {
		s.PrevSettlePrice = val
	}
	if val, ok := m["DECIMALS"].(int); ok {
		s.Decimals = val
	}
	if val, ok := m["MINSTEP"].(float64); ok {
		s.MinStep = val
	}
	if val, ok := m["LASTTRADEDATE"].(string); ok {
		s.LastTradeDate = val
	}
	if val, ok := m["LASTDELDATE"].(string); ok {
		s.LastDelDate = val
	}
	if val, ok := m["SECTYPE"].(string); ok {
		s.SecType = val
	}
	if val, ok := m["LATNAME"].(string); ok {
		s.LatName = val
	}
	if val, ok := m["ASSETCODE"].(string); ok {
		s.AssetCode = val
	}
	if val, ok := m["PREVOPENPOSITION"].(int); ok {
		s.PreviosOpenPosition = val
	}
	if val, ok := m["LOTVOLUME"].(int); ok {
		s.LotVolume = val
	}
	if val, ok := m["INITIALMARGIN"].(float64); ok {
		s.InitalMargin = val
	}
	if val, ok := m["HIGHLIMIT"].(float64); ok {
		s.HighLimit = val
	}
	if val, ok := m["LOWLIMIT"].(float64); ok {
		s.LowLimit = val
	}
	if val, ok := m["STEPPRICE"].(float64); ok {
		s.StepPrice = val
	}
	if val, ok := m["LASTSETTLPRICE"].(float64); ok {
		s.LastSettlPrice = val
	}
	if val, ok := m["PREVPRICE"].(float64); ok {
		s.PrevPrice = val
	}
	if val, ok := m["IMTIME"].(string); ok {
		s.ImTime = val
	}
	if val, ok := m["BUYSELLFEE"].(float64); ok {
		s.BuySellFee = val
	}
	if val, ok := m["SCALPERFEE"].(float64); ok {
		s.ScalperFee = val
	}
	if val, ok := m["NEGOTIATEDFEE"].(float64); ok {
		s.NegotiatedFee = val
	}
	if val, ok := m["EXERCISEFEE"].(float64); ok {
		s.ExerciseFee = val
	}

	return s
}

type MarketData struct {
	SecId                 string  `json:"sec_id"`                    // Идентификатор ценной бумаги.
	BoardId               string  `json:"board_id"`                  // Идентификатор рынка.
	Bid                   float64 `json:"bid"`                       // Цена предложения (покупки).
	Offer                 float64 `json:"offer"`                     // Цена предложения (продажи).
	Spread                float64 `json:"spread"`                    // Разница между ценами предложения и спроса.
	Open                  float64 `json:"open"`                      // Цена открытия.
	High                  float64 `json:"high"`                      // Максимальная цена.
	Low                   float64 `json:"low"`                       // Минимальная цена.
	Last                  float64 `json:"last"`                      // Цена последней сделки.
	Quantity              int     `json:"quantity"`                  // Количество.
	LastChange            float64 `json:"last_change"`               // Изменение цены.
	SettlePrice           float64 `json:"settle_price"`              // Цена закрытия.
	SettleToPrevSettle    float64 `json:"settle_to_prev_settle"`     // Разница между ценой закрытия и предыдущей ценой закрытия.
	OpenPosition          float64 `json:"open_position"`             // Открытая позиция.
	NumTrades             int     `json:"num_trades"`                // Количество сделок.
	VolToday              int     `json:"vol_today"`                 // Объем сделок.
	ValToday              int     `json:"val_today"`                 // Объем сделок в валюте рынка.
	ValTodayUsd           int     `json:"val_today_usd"`             // Объем сделок в долларах.
	UpdateTime            string  `json:"update_time"`               // Время обновления.
	LastChangePrcnt       float64 `json:"last_change_prcnt"`         // Изменение цены в процентах.
	BidDepth              int     `json:"bid_depth"`                 // Глубина спроса.
	BidDepthT             int     `json:"bid_depth_t"`               // Глубина спроса в тиках.
	NumBids               int     `json:"num_bids"`                  // Количество предложений.
	OfferDepth            int     `json:"offer_depth"`               // Глубина предложения.
	OfferDepthT           int     `json:"offer_depth_t"`             // Глубина предложения в тиках.
	NumOffers             int     `json:"num_offers"`                // Количество предложений на продажу.
	Time                  string  `json:"time"`                      // Время.
	SettleToPrevSettlePrc float64 `json:"settle_to_prev_settle_prc"` // Процентное изменение цены закрытия.
	SeqNum                int     `json:"seq_num"`                   // Уникальный номер последовательности.
	SysTime               string  `json:"sys_time"`                  // Время системы.
	TradeDate             string  `json:"trade_date"`                // Дата сделки.
	LastToPrevPrice       float64 `json:"last_to_prev_price"`        // Процентное изменение последней цены от предыдущей цены.
	OiChange              int     `json:"oi_change"`                 // Изменение открытых позиций.
	OpenPeriodPrice       float64 `json:"open_period_price"`         // Цена открытия за период.
	SwapRate              float64 `json:"swap_rate"`                 // Ставка свопа.
}

func (r *Response) MapMarketDataSingleResponse() MarketData {
	data := utils.MapData(r.MarketData.Columns, r.MarketData.Data)
	m := data[0]
	return mapMarketData(m)
}

func (r *Response) MapMarketDataFullResponse() []MarketData {
	res := []MarketData{}
	data := utils.MapData(r.MarketData.Columns, r.MarketData.Data)
	for _, m := range data {
		res = append(res, mapMarketData(m))
	}
	return res
}

func mapMarketData(m map[string]interface{}) MarketData {
	md := MarketData{}

	if val, ok := m["SECID"].(string); ok {
		md.SecId = val
	}
	if val, ok := m["BOARDID"].(string); ok {
		md.BoardId = val
	}
	if val, ok := m["BID"].(float64); ok {
		md.Bid = val
	}
	if val, ok := m["OFFER"].(float64); ok {
		md.Offer = val
	}
	if val, ok := m["SPREAD"].(float64); ok {
		md.Spread = val
	}
	if val, ok := m["OPEN"].(float64); ok {
		md.Open = val
	}
	if val, ok := m["HIGH"].(float64); ok {
		md.High = val
	}
	if val, ok := m["LOW"].(float64); ok {
		md.Low = val
	}
	if val, ok := m["LAST"].(float64); ok {
		md.Last = val
	}
	if val, ok := m["QUANTITY"].(int); ok {
		md.Quantity = val
	}
	if val, ok := m["LASTCHANGE"].(float64); ok {
		md.LastChange = val
	}
	if val, ok := m["SETTLEPRICE"].(float64); ok {
		md.SettlePrice = val
	}
	if val, ok := m["SETTLETOPREVSETTLE"].(float64); ok {
		md.SettleToPrevSettle = val
	}
	if val, ok := m["OPENPOSITION"].(float64); ok {
		md.OpenPosition = val
	}
	if val, ok := m["NUMTRADES"].(int); ok {
		md.NumTrades = val
	}
	if val, ok := m["VOLTODAY"].(int); ok {
		md.VolToday = val
	}
	if val, ok := m["VALTODAY"].(int); ok {
		md.ValToday = val
	}
	if val, ok := m["VALTODAY_USD"].(int); ok {
		md.ValTodayUsd = val
	}
	if val, ok := m["UPDATETIME"].(string); ok {
		md.UpdateTime = val
	}
	if val, ok := m["LASTCHANGEPRCNT"].(float64); ok {
		md.LastChangePrcnt = val
	}
	if val, ok := m["BIDDEPTH"].(int); ok {
		md.BidDepth = val
	}
	if val, ok := m["BIDDEPTHT"].(int); ok {
		md.BidDepthT = val
	}
	if val, ok := m["NUMBIDS"].(int); ok {
		md.NumBids = val
	}
	if val, ok := m["OFFERDEPTH"].(int); ok {
		md.OfferDepth = val
	}
	if val, ok := m["OFFERDEPTHT"].(int); ok {
		md.OfferDepthT = val
	}
	if val, ok := m["NUMOFFERS"].(int); ok {
		md.NumOffers = val
	}
	if val, ok := m["TIME"].(string); ok {
		md.Time = val
	}
	if val, ok := m["SETTLETOPREVSETTLEPRC"].(float64); ok {
		md.SettleToPrevSettlePrc = val
	}
	if val, ok := m["SEQNUM"].(int); ok {
		md.SeqNum = val
	}
	if val, ok := m["SYSTIME"].(string); ok {
		md.SysTime = val
	}
	if val, ok := m["TRADEDATE"].(string); ok {
		md.TradeDate = val
	}
	if val, ok := m["LASTTOPREVPRICE"].(float64); ok {
		md.LastToPrevPrice = val
	}
	if val, ok := m["OICHANGE"].(int); ok {
		md.OiChange = val
	}
	if val, ok := m["OPENPERIODPRICE"].(float64); ok {
		md.OpenPeriodPrice = val
	}
	if val, ok := m["SWAPRATE"].(float64); ok {
		md.SwapRate = val
	}

	return md
}

type Candle struct {
	Open   float64 `json:"open"`   // Цена открытия.
	Close  float64 `json:"close"`  // Цена закрытия.
	High   float64 `json:"high"`   // Максимальная цена.
	Low    float64 `json:"low"`    // Минимальная цена.
	Value  int     `json:"value"`  // Объем в рублях
	Volume float64 `json:"volume"` // Объем в лотах
	Begin  string  `json:"begin"`  // Начало свечки
	End    string  `json:"end"`    // Конец свечки
}

func (r *Response) MapCandleResponse() []Candle {
	res := []Candle{}
	data := utils.MapData(r.Candles.Columns, r.Candles.Data)
	for _, m := range data {
		res = append(res, mapCandleData(m))
	}
	return res
}

func mapCandleData(m map[string]interface{}) Candle {
	c := Candle{}

	if val, ok := m["open"].(float64); ok {
		c.Open = val
	}
	if val, ok := m["close"].(float64); ok {
		c.Close = val
	}
	if val, ok := m["high"].(float64); ok {
		c.High = val
	}
	if val, ok := m["low"].(float64); ok {
		c.Low = val
	}
	if val, ok := m["value"].(int); ok {
		c.Value = val
	}
	if val, ok := m["volume"].(float64); ok {
		c.Volume = val
	}
	if val, ok := m["begin"].(string); ok {
		c.Begin = val
	}
	if val, ok := m["end"].(string); ok {
		c.End = val
	}

	return c
}

type OrderBook struct {
	BoardId    string  `json:"board_id"`    // Идентификатор рынка.
	SecId      string  `json:"sec_id"`      // Идентификатор ценной бумаги.
	BuySell    string  `json:"buy_sell"`    // Индикатор покупки (B) или продажи (S).
	Price      float64 `json:"price"`       // Цена предложения или спроса.
	Quantity   int     `json:"quantity"`    // Количество товара, предлагаемого или запрошенного.
	SeqNum     int     `json:"seq_num"`     // Уникальный номер последовательности.
	UpdateTime string  `json:"update_time"` // Время обновления информации.
	Decimals   int     `json:"decimals"`    // Количество десятичных знаков для округления цены.
}

func (r *Response) MapOrderBookResponse() []OrderBook {
	res := []OrderBook{}
	data := utils.MapData(r.OrderBooks.Columns, r.OrderBooks.Data)
	for _, m := range data {
		res = append(res, mapOrderBookData(m))
	}
	return res
}

func mapOrderBookData(m map[string]interface{}) OrderBook {
	ob := OrderBook{}

	if val, ok := m["BOARDID"].(string); ok {
		ob.BoardId = val
	}
	if val, ok := m["SECID"].(string); ok {
		ob.SecId = val
	}
	if val, ok := m["BUYSELL"].(string); ok {
		ob.BuySell = val
	}
	if val, ok := m["PRICE"].(float64); ok {
		ob.Price = val
	}
	if val, ok := m["QUANTITY"].(int); ok {
		ob.Quantity = val
	}
	if val, ok := m["SEQNUM"].(int); ok {
		ob.SeqNum = val
	}
	if val, ok := m["UPDATETIME"].(string); ok {
		ob.UpdateTime = val
	}
	if val, ok := m["DECIMALS"].(int); ok {
		ob.Decimals = val
	}

	return ob
}

type Trade struct {
	TradeNo       int     `json:"trade_no"`        // Уникальный идентификатор сделки.
	BoardName     string  `json:"board_name"`      // Название рынка.
	SecId         string  `json:"sec_id"`          // Идентификатор ценной бумаги.
	TradeDate     string  `json:"trade_date"`      // Дата сделки.
	TradeTime     string  `json:"trade_time"`      // Время сделки.
	Price         float64 `json:"price"`           // Цена сделки.
	Quantity      int     `json:"quantity"`        // Количество товара, проданного или купленного в сделке.
	SysTime       string  `json:"sys_time"`        // Дата и время записи сделки в систему.
	RecNo         int     `json:"rec_no"`          // Уникальный номер записи.
	OpenPosition  int     `json:"open_position"`   // Открытая позиция.
	OffMarketDeal int     `json:"off_market_deal"` // Признак внебиржевой сделки (Оффмаркетная сделка).
	BuySell       string  `json:"buy_sell"`        // Индикатор покупки (B) или продажи (S).
}

func (r *Response) MapTradeResponse() []Trade {
	res := []Trade{}
	data := utils.MapData(r.Trades.Columns, r.Trades.Data)
	for _, m := range data {
		res = append(res, mapTradeData(m))
	}
	return res
}

func mapTradeData(m map[string]interface{}) Trade {
	t := Trade{}

	if val, ok := m["TRADENO"].(int); ok {
		t.TradeNo = val
	}
	if val, ok := m["BOARDNAME"].(string); ok {
		t.BoardName = val
	}
	if val, ok := m["SECID"].(string); ok {
		t.SecId = val
	}
	if val, ok := m["TRADEDATE"].(string); ok {
		t.TradeDate = val
	}
	if val, ok := m["TRADETIME"].(string); ok {
		t.TradeTime = val
	}
	if val, ok := m["PRICE"].(float64); ok {
		t.Price = val
	}
	if val, ok := m["QUANTITY"].(int); ok {
		t.Quantity = val
	}
	if val, ok := m["SYSTIME"].(string); ok {
		t.SysTime = val
	}
	if val, ok := m["RECNO"].(int); ok {
		t.RecNo = val
	}
	if val, ok := m["OPENPOSITION"].(int); ok {
		t.OpenPosition = val
	}
	if val, ok := m["OFFMARKETDEAL"].(int); ok {
		t.OffMarketDeal = val
	}
	if val, ok := m["BUYSELL"].(string); ok {
		t.BuySell = val
	}

	return t
}

type TradeStats struct {
	TradeDate  string  `json:"trade_date"`   // дата сделки
	TradeTime  string  `json:"trade_time"`   // время сделки
	SecId      string  `json:"sec_id"`       // код инструмента
	AssetCode  string  `json:"asset_code"`   // Код базового актива
	PrOpen     float64 `json:"pr_open"`      // цена открытия
	PrHigh     float64 `json:"pr_high"`      // максимальная цена
	PrLow      float64 `json:"pr_low"`       // минимальная цена
	PrClose    float64 `json:"pr_close"`     // цена закрытия
	PrStd      float64 `json:"pr_std"`       // стандартное отклонение цены
	Vol        int     `json:"vol"`          // объем в контрактах
	Val        int     `json:"val"`          // оборот в рублях
	Trades     int     `json:"trades"`       // кол-во сделок
	PrVwap     float64 `json:"pr_vwap"`      // средневзвешенная цена
	PrChange   float64 `json:"pr_change"`    // изменение цены в базисных пунктах
	TradesB    int     `json:"frades_b"`     // кол-во сделок на покупку
	TradesS    int     `json:"frades_s"`     // кол-во сделок на продажу
	ValB       float64 `json:"val_b"`        // оборот в рублях на покупку
	ValS       float64 `json:"val_s"`        // оборот в рублях на продажу
	VolB       int     `json:"vol_b"`        // объем на покупку в лотах
	VolS       int     `json:"vol_s"`        // объем на продажу в лотах
	Disb       float64 `json:"disb"`         // дисбаланс объема
	PrVwapB    float64 `json:"pr_vwap_b"`    // средневзвешенная цена на покупку
	PrVwapS    float64 `json:"pr_vwap_s"`    // средневзвешенная цена на продажу
	Im         float64 `json:"im"`           // гарантийное обеспечение
	OiOpen     int     `json:"oi_open"`      // ОИ на открытии
	OiHigh     int     `json:"oi_high"`      // максимальный ОИ
	OiLow      int     `json:"oi_low"`       // минимальный ОИ
	OiClose    int     `json:"oi_close"`     // ОИ на закрытии
	SecPrOpen  int     `json:"sec_pr_open"`  // кол-во секунд до pr_open
	SecPrHigh  int     `json:"sec_pr_high"`  // кол-во секунд до pr_high
	SecPrLow   int     `json:"sec_pr_low"`   // кол-во секунд до pr_low
	SecPrClose int     `json:"sec_pr_close"` // кол-во секунд до pr_close
	SysTime    string  `json:"SYSTIME"`      // время системы
}

func (r *Response) TradeStatsResponse() []TradeStats {
	res := []TradeStats{}
	data := utils.MapData(r.Data.Columns, r.Data.Data)
	for _, m := range data {
		res = append(res, mapTradeStats(m))
	}
	return res
}

func mapTradeStats(m map[string]interface{}) TradeStats {
	t := TradeStats{}

	if val, ok := m["tradedate"].(string); ok {
		t.TradeDate = val
	}
	if val, ok := m["tradetime"].(string); ok {
		t.TradeTime = val
	}
	if val, ok := m["secid"].(string); ok {
		t.SecId = val
	}
	if val, ok := m["asset_code"].(string); ok {
		t.AssetCode = val
	}
	if val, ok := m["pr_open"].(float64); ok {
		t.PrOpen = val
	}
	if val, ok := m["pr_high"].(float64); ok {
		t.PrHigh = val
	}
	if val, ok := m["pr_low"].(float64); ok {
		t.PrLow = val
	}
	if val, ok := m["pr_close"].(float64); ok {
		t.PrClose = val
	}
	if val, ok := m["pr_std"].(float64); ok {
		t.PrStd = val
	}
	if val, ok := m["vol"].(int); ok {
		t.Vol = val
	}
	if val, ok := m["val"].(int); ok {
		t.Val = val
	}
	if val, ok := m["trades"].(int); ok {
		t.Trades = val
	}
	if val, ok := m["pr_vwap"].(float64); ok {
		t.PrVwap = val
	}
	if val, ok := m["pr_change"].(float64); ok {
		t.PrChange = val
	}
	if val, ok := m["trades_b"].(int); ok {
		t.TradesB = val
	}
	if val, ok := m["trades_s"].(int); ok {
		t.TradesS = val
	}
	if val, ok := m["val_b"].(float64); ok {
		t.ValB = val
	}
	if val, ok := m["val_s"].(float64); ok {
		t.ValS = val
	}
	if val, ok := m["vol_b"].(int); ok {
		t.VolB = val
	}
	if val, ok := m["vol_s"].(int); ok {
		t.VolS = val
	}
	if val, ok := m["disb"].(float64); ok {
		t.Disb = val
	}
	if val, ok := m["pr_vwap_b"].(float64); ok {
		t.PrVwapB = val
	}
	if val, ok := m["pr_vwap_s"].(float64); ok {
		t.PrVwapS = val
	}
	if val, ok := m["im"].(float64); ok {
		t.Im = val
	}
	if val, ok := m["oi_open"].(int); ok {
		t.OiOpen = val
	}
	if val, ok := m["oi_high"].(int); ok {
		t.OiHigh = val
	}
	if val, ok := m["oi_low"].(int); ok {
		t.OiLow = val
	}
	if val, ok := m["oi_close"].(int); ok {
		t.OiClose = val
	}
	if val, ok := m["sec_pr_open"].(int); ok {
		t.SecPrOpen = val
	}
	if val, ok := m["sec_pr_high"].(int); ok {
		t.SecPrHigh = val
	}
	if val, ok := m["sec_pr_low"].(int); ok {
		t.SecPrLow = val
	}
	if val, ok := m["sec_pr_close"].(int); ok {
		t.SecPrClose = val
	}
	if val, ok := m["SYSTIME"].(string); ok {
		t.SysTime = val
	}

	return t
}

type ObStats struct {
	TradeDate  string  `json:"trade_date"`  // дата сделки
	TradeTime  string  `json:"trade_time"`  // время сделки
	SecId      string  `json:"sec_id"`      // код инструмента
	AssetCode  string  `json:"asset_code"`  // Код базового актива
	MidPrice   float64 `json:"mid_price"`   // средняя цена между лучшей ценой на продажи и лучшей ценой на покупку
	MicroPrice float64 `json:"micro_price"` // средневзвешенная цена между лучшей ценой на продажи и лучшей ценой на покупку
	SpreadL1   float64 `json:"spread_l1"`   // спред цены в базисных пунктах
	SpreadL2   float64 `json:"spread_l2"`   // спред между 2ым уровнем цен покупки и продажи
	SpreadL3   float64 `json:"spread_l3"`   // спред между 3ым уровнем цен покупки и продажи
	SpreadL5   float64 `json:"spread_l5"`   // спред между 5ым уровнем цен покупки и продажи
	SpreadL10  float64 `json:"spread_l10"`  // спред между 10ым уровнем цен покупки и продажи
	SpreadL20  float64 `json:"spread_l20"`  // спред между 20ым уровнем цен покупки и продажи
	LevelsB    int     `json:"levels_b"`    // кол-во уровней цен в стакане (покупка)
	LevelsS    int     `json:"levels_s"`    // кол-во уровней цен в стакане (продажа)
	VolBL1     int     `json:"vol_b_l1"`    // объем на первом уровне (покупка)
	VolBL2     int     `json:"vol_b_l2"`    // объем на первых двух уровнях (покупка)
	VolBL3     int     `json:"vol_b_l3"`    // объем на первых трех уровнях (покупка)
	VolBL5     int     `json:"vol_b_l5"`    // объем на первых пяти уровнях (покупка)
	VolBL10    int     `json:"vol_b_l10"`   // объем на первых десяти уровнях (покупка)
	VolBL20    int     `json:"vol_b_l20"`   // объем на первых двадцати уровнях (покупка)
	VolSL1     int     `json:"vol_s_l1"`    // объем на первом уровне (продажа)
	VolSL2     int     `json:"vol_s_l2"`    // объем на первых двух уровнях (продажа)
	VolSL3     int     `json:"vol_s_l3"`    // объем на первых трех уровнях (продажа)
	VolSL5     int     `json:"vol_s_l5"`    // объем на первых пяти уровнях (продажа)
	VolSL10    int     `json:"vol_s_l10"`   // объем на первых десяти уровнях (продажа)
	VolSL20    int     `json:"vol_s_l20"`   // объем на первых двадцати уровнях (продажа)
	VwapBL3    float64 `json:"vwap_b_l3"`   // средневзвешенная цена по первым трем уровням на покупку
	VwapBL5    float64 `json:"vwap_b_l5"`   // средневзвешенная цена по первым пяти уровням на покупку
	VwapBL10   float64 `json:"vwap_b_l10"`  // средневзвешенная цена по первым десяти уровням на покупку
	VwapBL20   float64 `json:"vwap_b_l20"`  // средневзвешенная цена по первым двадцати уровням на покупку
	VwapSL3    float64 `json:"vwap_s_l3"`   // средневзвешенная цена по первым трем уровням на продажу
	VwapSL5    float64 `json:"vwap_s_l5"`   // средневзвешенная цена по первым пяти уровням на продажу
	VwapSL10   float64 `json:"vwap_s_l10"`  // средневзвешенная цена по первым десяти уровням на продажу
	VwapSL20   float64 `json:"vwap_s_l20"`  // средневзвешенная цена по первым двадцати уровням на продажу
	SysTime    string  `json:"SYSTIME"`     // время системы
}

func (r *Response) ObStatsResponse() []ObStats {
	res := []ObStats{}
	data := utils.MapData(r.Data.Columns, r.Data.Data)
	for _, m := range data {
		res = append(res, mapObStats(m))
	}
	return res
}

func mapObStats(m map[string]interface{}) ObStats {
	t := ObStats{}

	if val, ok := m["tradedate"].(string); ok {
		t.TradeDate = val
	}
	if val, ok := m["tradetime"].(string); ok {
		t.TradeTime = val
	}
	if val, ok := m["secid"].(string); ok {
		t.SecId = val
	}
	if val, ok := m["asset_code"].(string); ok {
		t.AssetCode = val
	}
	if val, ok := m["mid_price"].(float64); ok {
		t.MidPrice = val
	}
	if val, ok := m["micro_price"].(float64); ok {
		t.MicroPrice = val
	}
	if val, ok := m["spread_l1"].(float64); ok {
		t.SpreadL1 = val
	}
	if val, ok := m["spread_l2"].(float64); ok {
		t.SpreadL2 = val
	}
	if val, ok := m["spread_l3"].(float64); ok {
		t.SpreadL3 = val
	}
	if val, ok := m["spread_l5"].(float64); ok {
		t.SpreadL5 = val
	}
	if val, ok := m["spread_l10"].(float64); ok {
		t.SpreadL10 = val
	}
	if val, ok := m["spread_l20"].(float64); ok {
		t.SpreadL20 = val
	}
	if val, ok := m["levels_b"].(int); ok {
		t.LevelsB = val
	}
	if val, ok := m["levels_s"].(int); ok {
		t.LevelsS = val
	}
	if val, ok := m["vol_b_l1"].(int); ok {
		t.VolBL1 = val
	}
	if val, ok := m["vol_b_l2"].(int); ok {
		t.VolBL2 = val
	}
	if val, ok := m["vol_b_l3"].(int); ok {
		t.VolBL3 = val
	}
	if val, ok := m["vol_b_l5"].(int); ok {
		t.VolBL5 = val
	}
	if val, ok := m["vol_b_l10"].(int); ok {
		t.VolBL10 = val
	}
	if val, ok := m["vol_b_l20"].(int); ok {
		t.VolBL20 = val
	}
	if val, ok := m["vol_s_l1"].(int); ok {
		t.VolSL1 = val
	}
	if val, ok := m["vol_s_l2"].(int); ok {
		t.VolSL2 = val
	}
	if val, ok := m["vol_s_l3"].(int); ok {
		t.VolSL3 = val
	}
	if val, ok := m["vol_s_l5"].(int); ok {
		t.VolSL5 = val
	}
	if val, ok := m["vol_s_l10"].(int); ok {
		t.VolSL10 = val
	}
	if val, ok := m["vol_s_l20"].(int); ok {
		t.VolSL20 = val
	}
	if val, ok := m["vwap_b_l3"].(float64); ok {
		t.VwapBL3 = val
	}
	if val, ok := m["vwap_b_l5"].(float64); ok {
		t.VwapBL5 = val
	}
	if val, ok := m["vwap_b_l10"].(float64); ok {
		t.VwapBL10 = val
	}
	if val, ok := m["vwap_b_l20"].(float64); ok {
		t.VwapBL20 = val
	}
	if val, ok := m["vwap_s_l3"].(float64); ok {
		t.VwapSL3 = val
	}
	if val, ok := m["vwap_s_l5"].(float64); ok {
		t.VwapSL5 = val
	}
	if val, ok := m["vwap_s_l10"].(float64); ok {
		t.VwapSL10 = val
	}
	if val, ok := m["vwap_s_l20"].(float64); ok {
		t.VwapSL20 = val
	}
	if val, ok := m["SYSTIME"].(string); ok {
		t.SysTime = val
	}

	return t
}
