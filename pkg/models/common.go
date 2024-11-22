package models

import "github.com/DBoyara/MoexAlgoPackGo/pkg/utils"

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
