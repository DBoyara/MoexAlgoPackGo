package models

import (
	"github.com/DBoyara/MoexAlgoPackGo/pkg/utils"
)

type Response struct {
	Securities Data `json:"securities,omitempty"`
	MarketData Data `json:"marketdata,omitempty"`
	Candles    Data `json:"candles,omitempty"`
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
	data := utils.MapData(r.Securities.Columns, r.MarketData.Data)
	m := data[0]
	return mapMarketData(m)
}

func (r *Response) MapMarketDataFullResponse() []MarketData {
	res := []MarketData{}
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
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
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
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
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
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
	data := utils.MapData(r.Securities.Columns, r.Securities.Data)
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
