package network

import (
	"fmt"
	"teletraffic/backend/traffic"
)

func Simple_BSC(bhca, sms, reg, callLoad float64) float64 {
	result := (1800 / callLoad) * (1 / (bhca + (2 * sms / 3) + (reg / 3)))
	return result
}

func Basic_traffic(bhca, mht float64) float64 {
	traff := (bhca / 3600) * mht * 1000
	return traffic.Round(traff, 3)
}

func TRX(subs, subs_per_erl, trx_per_erl float64) float64 {
	trx := (subs * subs_per_erl) / trx_per_erl
	return traffic.Round(trx, 3)
}

func Cells(trx, trx_per_erl float64) float64 {
	fmt.Println(trx, trx_per_erl)
	return traffic.Round((trx / trx_per_erl), 3)
}
