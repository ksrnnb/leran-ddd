package main

import (
	"log"

	"github.com/ksrnnb/learn-ddd/route"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.RegisterRoute(e)

	log.Fatal(e.Start(":3000"))
}
