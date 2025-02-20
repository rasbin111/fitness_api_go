package main 

import (
    "github.com/labstack/echo/v4"
    "fitness-api/cmd/handlers"
    "fitness-api/cmd/storage"
)

func main(){
    e := echo.New()
    e.GET("/", handlers.Home)
    storage.InitDB()
    e.Logger.Fatal(e.Start(":8080"))
}
