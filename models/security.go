package models

type Security struct {
	SecId         string  `json:"sec_id"`
	ShortName     string  `json:"short_name"`
	LastTradeDate string  `json:"last_trade_date"`
	Lot           float64 `json:"lot"`
	Margin        float64 `json:"margin"`
}
