package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"teletraffic/backend/network"
)

func Calc_simple_BSC(c echo.Context) error {
	bhca := c.Param("bhca")
	sms := c.Param("sms")
	reg := c.Param("reg")
	callLoad := c.Param("callLoad")

	bhcaCvt, _ := strconv.ParseFloat(bhca, 64)
	smsCvt, _ := strconv.ParseFloat(sms, 64)
	regCvt, _ := strconv.ParseFloat(reg, 64)
	callLoadCvt, _ := strconv.ParseFloat(callLoad, 64)

	res := network.Simple_BSC(bhcaCvt, smsCvt, regCvt, callLoadCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_basic_traffic(c echo.Context) error {
	bhca := c.Param("bhca")
	mht := c.Param("mht")
	bhcaCvt, _ := strconv.ParseFloat(bhca, 64)
	mhtCvt, _ := strconv.ParseFloat(mht, 64)

	res := network.Basic_traffic(bhcaCvt, mhtCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_TRX(c echo.Context) error {
	subs := c.Param("subs")
	subs_per_erl := c.Param("subs_per_erl")
	trx_erl := c.Param("trx_erl")

	subsCvt, _ := strconv.ParseFloat(subs, 64)
	subsErlCvt, _ := strconv.ParseFloat(subs_per_erl, 64)
	trxErlCvt, _ := strconv.ParseFloat(trx_erl, 64)

	res := network.TRX(subsCvt, subsErlCvt, trxErlCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_Cells(c echo.Context) error {
	trx := c.Param("trx")
	trx_cell := c.Param("trx_cell")
	
	trxCvt, _ := strconv.ParseFloat(trx, 64)
	trxErlCvt, _ := strconv.ParseFloat(trx_cell, 64)

	res := network.Cells(trxCvt, trxErlCvt)

	return c.JSON(http.StatusOK, res)

}
