package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"teletraffic/backend/traffic"
)

/*==========================================================*/
/*========================First Row========================*/
/*==========================================================*/

var erlangB traffic.ErlangB
var erlangC traffic.ErlangC
var poisson traffic.Poisson
var calcs traffic.Calculations

func Calc_traffic_handler(c echo.Context) error {
	gos := c.Param("gos")
	gosCvt, _ := strconv.ParseFloat(gos, 64)

	num := c.Param("num_of_channels")
	numCvt, _ := strconv.Atoi(num)

	pathSegments := strings.Split(c.Path(), "/")
	erlangType := pathSegments[2]

	switch erlangType {
	case "erlangB":
		traff := traffic.ErlangB.Calc_traffic(erlangB, numCvt, gosCvt)
		return c.JSON(http.StatusOK, traff)
	case "erlangC":
		traff := traffic.ErlangC.Calc_traffic(erlangC, numCvt, gosCvt)
		return c.JSON(http.StatusOK, traff)
	case "poisson":
		traff := traffic.Poisson.Calc_traffic(poisson, numCvt, gosCvt)
		return c.JSON(http.StatusOK, traff)
	default:
		return fmt.Errorf("ErlangType isn't recognized")
	}
}

func Calc_gos_handler(c echo.Context) error {
	traff := c.Param("traffic")
	traffCvt, _ := strconv.ParseFloat(traff, 64)

	num := c.Param("num_of_channels")
	numCvt, _ := strconv.Atoi(num)
	pathSegments := strings.Split(c.Path(), "/")
	erlangType := pathSegments[2]

	switch erlangType {
	case "erlangB":
		gos := traffic.ErlangB.Calc_gos(erlangB, numCvt, traffCvt)
		return c.JSON(http.StatusOK, gos)
	case "erlangC":
		gos := traffic.ErlangC.Calc_gos(erlangC, numCvt, traffCvt)
		return c.JSON(http.StatusOK, gos)
	case "poisson":
		gos := traffic.Poisson.Calc_gos(poisson, numCvt, traffCvt)
		return c.JSON(http.StatusOK, gos)
	default:
		return fmt.Errorf("ErlangType isn't recognized")
	}
}

func Calc_num_of_channels_handler(c echo.Context) error {

	gos := c.Param("gos")
	gosCvt, _ := strconv.ParseFloat(gos, 64)

	traff := c.Param("traffic")
	traffCvt, _ := strconv.ParseFloat(traff, 64)

	pathSegments := strings.Split(c.Path(), "/")
	erlangType := pathSegments[2]

	switch erlangType {
	case "erlangB":
		num := traffic.ErlangB.Calc_num_channels(erlangB, traffCvt, gosCvt)
		return c.JSON(http.StatusOK, num)
	case "erlangC":
		num := traffic.ErlangC.Calc_num_channels(erlangC, traffCvt, gosCvt)
		return c.JSON(http.StatusOK, num)
	case "poisson":
		num := traffic.Poisson.Calc_num_channels(poisson, traffCvt, gosCvt)
		return c.JSON(http.StatusOK, num)
	default:
		return fmt.Errorf("ErlangType isn't recognized")
	}
}

/*==========================================================*/
/*========================Second Row========================*/
/*==========================================================*/

func Calc_num_of_subs(c echo.Context) error {
	traff := c.Param("traffic_2")
	traffCvt, _ := strconv.ParseFloat(traff, 64)

	traffic_per_subs := c.Param("traffic_per_subs")
	traffic_per_subs_cvt, _ := strconv.ParseFloat(traffic_per_subs, 64)

	num_of_subs := traffic.Calculations.Calc_num_of_subs(calcs, traffCvt, traffic_per_subs_cvt)

	return c.JSON(http.StatusOK, num_of_subs)
}

func Calc_traffic_per_subs(c echo.Context) error {
	traff := c.Param("traffic_2")
	traffCvt, _ := strconv.ParseFloat(traff, 64)

	num_of_subs := c.Param("num_of_subs")
	num_of_subs_cvt, _ := strconv.Atoi(num_of_subs)

	traffic_per_subs := traffic.Calculations.Calc_traffic_per_subs(calcs, traffCvt, uint(num_of_subs_cvt))
	return c.JSON(http.StatusOK, traffic_per_subs)
}

func Calc_traffic_2(c echo.Context) error {
	traffic_per_subs := c.Param("traffic_per_subs")
	traffic_per_subs_cvt, _ := strconv.ParseFloat(traffic_per_subs, 64)

	num_of_subs := c.Param("num_of_subs")
	num_of_subs_cvt, _ := strconv.ParseFloat(num_of_subs, 64)

	traff := traffic.Calculations.Calc_traffic(calcs, num_of_subs_cvt, traffic_per_subs_cvt)
	return c.JSON(http.StatusOK, traff)
}

/*==========================================================*/
/*========================Third Row========================*/
/*==========================================================*/
func Calc_traffic_3(c echo.Context) error {
	bhca := c.Param("bhca")
	bhcaCvt, _ := strconv.ParseFloat(bhca, 64)

	mht := c.Param("mht")
	mhtCvt, _ := strconv.ParseFloat(mht, 64)

	traff_3 := traffic.Calculations.Calc_traffic_2(calcs, bhcaCvt, mhtCvt)
	return c.JSON(http.StatusOK, traff_3)
}

func Calc_bhca(c echo.Context) error {
	traff_3 := c.Param("traffic_3")
	traffCvt, _ := strconv.ParseFloat(traff_3, 64)

	mht := c.Param("mht")
	mhtCvt, _ := strconv.ParseFloat(mht, 64)

	bhca := traffic.Calculations.Calc_BHCA(calcs, traffCvt, mhtCvt)
	return c.JSON(http.StatusOK, bhca)
}

func Calc_mht(c echo.Context) error {
	bhca := c.Param("bhca")
	bhcaCvt, _ := strconv.ParseFloat(bhca, 64)

	traff_3 := c.Param("traffic_3")
	traffCvt, _ := strconv.ParseFloat(traff_3, 64)

	mht := traffic.Calculations.Calc_MHT(calcs, traffCvt, bhcaCvt)
	return c.JSON(http.StatusOK, mht)
}

/*==========================================================*/
/*========================Fourth Row========================*/
/*==========================================================*/

func Calc_subs_vlr(c echo.Context) error {
	traffic_per_subs := c.Param("traffic_per_subs_2")
	traffic_per_subs_cvt, _ := strconv.ParseFloat(traffic_per_subs, 64)

	switch_traffic := c.Param("switch_traffic")
	switch_traffic_cvt, _ := strconv.ParseFloat(switch_traffic, 64)

	subs_vlr := traffic.Calculations.Calc_subs_vlr(calcs, traffic_per_subs_cvt, switch_traffic_cvt)
	return c.JSON(http.StatusOK, subs_vlr)
}

func Calc_traffic_per_subs_2(c echo.Context) error {
	subs_vlr := c.Param("subs_vlr")
	subs_vlrs_cvt, _ := strconv.ParseFloat(subs_vlr, 64)

	switch_traffic := c.Param("switch_traffic")
	switch_traffic_cvt, _ := strconv.ParseFloat(switch_traffic, 64)

	traffic_per_subs := traffic.Calculations.Calc_traffic_per_subs_2(calcs, subs_vlrs_cvt, switch_traffic_cvt)
	return c.JSON(http.StatusOK, traffic_per_subs)
}

func Calc_switch_traffic(c echo.Context) error {
	traffic_per_subs := c.Param("traffic_per_subs_2")
	traffic_per_subs_cvt, _ := strconv.ParseFloat(traffic_per_subs, 64)

	subs_vlr := c.Param("subs_vlr")
	subs_vlrs_cvt, _ := strconv.ParseFloat(subs_vlr, 64)

	switch_traffic := traffic.Calculations.Calc_switch_traffic(calcs, subs_vlrs_cvt, traffic_per_subs_cvt)
	return c.JSON(http.StatusOK, switch_traffic)
}
