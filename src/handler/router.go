package handler

import "github.com/labstack/echo"

func InitRouting(e *echo.Echo, geoHandler GeoHandler) {
	e.GET("/address", geoHandler.Search())
	e.GET("/address/access_logs", geoHandler.History())
}
