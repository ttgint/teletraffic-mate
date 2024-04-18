package routes

import (
	"github.com/labstack/echo/v4"
	"teletraffic/backend/handlers"
)

func NetworkRoutesInit(e *echo.Echo) {
	e.POST("/network/simple-bsc/:bhca/:sms/:reg/:callLoad", handlers.Calc_simple_BSC)
	e.POST("/network/basic-traffic/:bhca/:mht", handlers.Calc_basic_traffic)
	e.POST("/network/trx/:subs/:subs_per_erl/:trx_erl", handlers.Calc_TRX)
	e.POST("/network/trx/:trx/:trx_cell", handlers.Calc_Cells)

}
