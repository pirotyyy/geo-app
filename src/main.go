package main

import (
	"fmt"
	"go-geo-backend/src/handler"
	"go-geo-backend/src/injector"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("server start")
	// load .go_env
	loadEnv()
	geoHandler := injector.InjectGeoHandler()
	e := echo.New()
	handler.InitRouting(e, geoHandler)
	// set apihost localhost:8080
	e.Logger.Fatal(e.Start(":8080"))
}

func loadEnv() {
	err := godotenv.Load(".go_env")

	if err != nil {
		panic(err)
	}
}
