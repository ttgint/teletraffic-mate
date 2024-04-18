package routes

import (
	"github.com/labstack/echo/v4"
	"teletraffic/backend/handlers"
)

func RFRoutesInit(e *echo.Echo) {
	e.POST("rf/eirp/:power/:combiner/:feeder/:antenna", handlers.Calc_EIRP)
	e.POST("rf/power/:eirp/:combiner/:feeder/:antenna", handlers.Calc_Trans_Power)
	e.POST("rf/combiner/:eirp/:power/:feeder/:antenna", handlers.Calc_Combiner)
	e.POST("rf/feeder/:eirp/:power/:combiner/:antenna", handlers.Calc_Feeder)
	e.POST("rf/antenna/:eirp/:power/:combiner/:feeder", handlers.Calc_Antenna)

}
