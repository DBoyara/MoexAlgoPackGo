package models

import "github.com/DBoyara/MoexAlgoPackGo/pkg/utils"

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
