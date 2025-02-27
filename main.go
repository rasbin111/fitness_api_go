package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	storage.InitDB()

	e.GET("/", handlers.Home)

    e.GET("/users", handlers.HandleListUser)

	e.POST("/users", handlers.HandleCreateUser)

	e.POST("/measurements", handlers.HandleCreateMeasurement)

	e.PUT("/users/:id", handlers.HandleUpdateUser)

    e.PUT("/measurements/:id", handlers.HandleCreateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))

}
