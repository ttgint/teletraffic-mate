package handlers

import (
	"net/http"
	"strconv"
	"teletraffic/backend/RF"
	"github.com/labstack/echo/v4"
)

func Calc_EIRP(c echo.Context) error{
	power := c.Param("power")
	powerCvt, _ := strconv.ParseFloat(power, 64)

	combiner := c.Param("combiner")
	combinerCvt, _ := strconv.ParseFloat(combiner, 64)

	feeder := c.Param("feeder")
	feederCvt, _ := strconv.ParseFloat(feeder, 64)

	antenna := c.Param("antenna")
	antennaCvt, _ := strconv.ParseFloat(antenna, 64)

	res := RF.Eirp(powerCvt, combinerCvt, feederCvt, antennaCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_Trans_Power(c echo.Context) error{
	eirp := c.Param("eirp")
	eirpCvt, _ := strconv.ParseFloat(eirp, 64)

	combiner := c.Param("combiner")
	combinerCvt, _ := strconv.ParseFloat(combiner, 64)

	feeder := c.Param("feeder")
	feederCvt, _ := strconv.ParseFloat(feeder, 64)

	antenna := c.Param("antenna")
	antennaCvt, _ := strconv.ParseFloat(antenna, 64)

	res := RF.TransmittedPower(eirpCvt, combinerCvt, feederCvt, antennaCvt)

	return c.JSON(http.StatusOK, res)
}


func Calc_Combiner(c echo.Context) error{
	eirp := c.Param("eirp")
	eirpCvt, _ := strconv.ParseFloat(eirp, 64)

	power := c.Param("power")
	powerCvt, _ := strconv.ParseFloat(power, 64)

	feeder := c.Param("feeder")
	feederCvt, _ := strconv.ParseFloat(feeder, 64)

	antenna := c.Param("antenna")
	antennaCvt, _ := strconv.ParseFloat(antenna, 64)

	res := RF.CombinerLoss(eirpCvt, powerCvt, feederCvt, antennaCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_Feeder(c echo.Context) error{
	eirp := c.Param("eirp")
	eirpCvt, _ := strconv.ParseFloat(eirp, 64)

	power := c.Param("power")
	powerCvt, _ := strconv.ParseFloat(power, 64)

	combiner := c.Param("combiner")
	combinerCvt, _ := strconv.ParseFloat(combiner, 64)

	antenna := c.Param("antenna")
	antennaCvt, _ := strconv.ParseFloat(antenna, 64)

	res := RF.FeederLoss(eirpCvt, powerCvt, combinerCvt, antennaCvt)

	return c.JSON(http.StatusOK, res)
}

func Calc_Antenna(c echo.Context) error{
	eirp := c.Param("eirp")
	eirpCvt, _ := strconv.ParseFloat(eirp, 64)

	power := c.Param("power")
	powerCvt, _ := strconv.ParseFloat(power, 64)

	combiner := c.Param("combiner")
	combinerCvt, _ := strconv.ParseFloat(combiner, 64)

	feeder := c.Param("feeder")
	feederCvt, _ := strconv.ParseFloat(feeder, 64)

	res := RF.AntennaGain(eirpCvt, powerCvt, combinerCvt, feederCvt)

	return c.JSON(http.StatusOK, res)
}