package main

import (
	"github.com/yekai1003/mswallet/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", router.Hello)
	e.POST("/register", router.Register)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
