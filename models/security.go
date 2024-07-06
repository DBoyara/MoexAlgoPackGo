package models

type SecuritiesResponse struct {
	Securities Securities `json:"securities"`
}

type Securities struct {
	Columns []string        `json:"columns"`
	Data    [][]interface{} `json:"data"`
}

func (sr *SecuritiesResponse) MapSecurity() Security {
	res := Security{}

	if len(sr.Securities.Data[0]) != 25 {
		return res
	}

	for _, row := range sr.Securities.Data {
		res.SecId = row[0].(string)
		res.ShortName = row[2].(string)
		res.LastTradeDate = row[7].(string)
		res.Lot = row[13].(float64)
		res.Margin = row[14].(float64)
	}
	return res
}

type Security struct {
	SecId         string  `json:"sec_id"`
	ShortName     string  `json:"short_name"`
	LastTradeDate string  `json:"last_trade_date"`
	Lot           float64 `json:"lot"`
	Margin        float64 `json:"margin"`
}
