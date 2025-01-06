package main

import (
	"github.com/WENDELLDELIMA/go-expert-multithreading/internal/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/cep/:cep", handlers.GetAddressHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
