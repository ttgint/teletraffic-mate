package routes

import (
	"github.com/labstack/echo/v4"
	"teletraffic/backend/handlers"
)

func TrafficRoutesInit(e *echo.Echo) {
	/***************************/
	/*Erlang Based Calculations*/
	/***************************/
	e.POST("/traffic/erlangB/traffic/:gos/:num_of_channels", handlers.Calc_traffic_handler)
	e.POST("/traffic/erlangB/gos/:traffic/:num_of_channels", handlers.Calc_gos_handler)
	e.POST("/traffic/erlangB/num-of-channels/:gos/:traffic", handlers.Calc_num_of_channels_handler)

	e.POST("/traffic/erlangC/traffic/:gos/:num_of_channels", handlers.Calc_traffic_handler)
	e.POST("/traffic/erlangC/gos/:traffic/:num_of_channels", handlers.Calc_gos_handler)
	e.POST("/traffic/erlangC/num-of-channels/:gos/:traffic", handlers.Calc_num_of_channels_handler)

	e.POST("/traffic/poisson/traffic/:gos/:num_of_channels", handlers.Calc_traffic_handler)
	e.POST("/traffic/poisson/gos/:traffic/:num_of_channels", handlers.Calc_gos_handler)
	e.POST("/traffic/poisson/num-of-channels/:gos/:traffic", handlers.Calc_num_of_channels_handler)
	/***************************/
	/*Erlang Based Calculations*/
	/***************************/

	/*Number Of Subs vs Traffic/Subs vs traffic*/
	e.POST("traffic/num-of-subs/:traffic_per_subs/:traffic_2", handlers.Calc_num_of_subs)
	e.POST("traffic/traffic-per-subs/:num_of_subs/:traffic_2", handlers.Calc_traffic_per_subs)
	e.POST("traffic/traffic-2/:num_of_subs/:traffic_per_subs", handlers.Calc_traffic_2)

	/*Traffic vs BHCA vs MHT*/
	e.POST("traffic/traffic-3/:bhca/:mht", handlers.Calc_traffic_3)
	e.POST("traffic/bhca/:traffic_3/:mht", handlers.Calc_bhca)
	e.POST("traffic/mht/:bhca/:traffic_3", handlers.Calc_mht)

	/*Subs/VLR vs Switch Traffic vs Traffic/Subs*/
	e.POST("traffic/subs-vlr/:traffic_per_subs_2/:switch_traffic", handlers.Calc_subs_vlr)
	e.POST("traffic/traffic-per-subs-2/:subs_vlr/:switch_traffic", handlers.Calc_traffic_per_subs_2)
	e.POST("traffic/switch-traffic/:traffic_per_subs_2/:subs_vlr", handlers.Calc_switch_traffic)

}
