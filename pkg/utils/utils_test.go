package utils

import (
	"reflect"
	"testing"
)

func TestMapData(t *testing.T) {
	keys := []string{
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
	}

	values := [][]interface{}{
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
	}

	expectedResult := []map[string]interface{}{
		{
			"SECID":            "MMM4",
			"BOARDID":          "RFUD",
			"SHORTNAME":        "MXI-6.24",
			"SECNAME":          "Фьючерсный контракт MXI-6.24",
			"PREVSETTLEPRICE":  3462.65,
			"DECIMALS":         2.0,
			"MINSTEP":          0.05000,
			"LASTTRADEDATE":    "2024-06-20",
			"LASTDELDATE":      "2024-06-20",
			"SECTYPE":          "MM",
			"LATNAME":          "MXI-6.24",
			"ASSETCODE":        "MXI",
			"PREVOPENPOSITION": 76868.0,
			"LOTVOLUME":        1.0,
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
		},
	}

	result := MapData(keys, values)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("MapData() = %v, want %v", result, expectedResult)
	}
}
