package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ksrnnb/learn-ddd/route"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	route.RegisterRoute(e)

	log.Fatal(e.Start(":3000"))
}
