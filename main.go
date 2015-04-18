package main

import "github.com/labstack/echo"
import "github.com/ifraixedes/twelveman/handlers"

func main() {
	e := echo.New()
	e.Get("/matches/:league", handlers.Matches)
	e.Get("/chants/:team", handlers.Chants)
	e.Run(":4000")
}
