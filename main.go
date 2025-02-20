package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	storage.InitDB()

	e.POST("/users", handlers.HandleCreateUser)

	e.POST("/measurements", handlers.HandleCreateMeasurement)

	e.PUT("/users/:id", handlers.HandleUpdateUser)

	e.Logger.Fatal(e.Start(":8080"))

}
