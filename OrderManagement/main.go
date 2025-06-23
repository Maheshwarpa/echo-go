package main

import (
	"log"
	"net/http"

	ord "github.com/maheshwarpa/echo-go/Order"
	pdt "github.com/maheshwarpa/echo-go/Product"

	//"github.com/labstack/echo"
	"github.com/labstack/echo/v4"
)

func init() {
	pdt.LoadData()
	ord.InitalizeOrderList()
}
func main() {

	e := echo.New()

	e.GET("/", Greetings)
	e.POST("/post", ord.CreateOrder)
	e.PUT("/update", ord.UpdateOrder)
	e.DELETE("/delete/:id", ord.DeleteOrder)

	log.Fatal(e.Start(":8080"))
}

func Greetings(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to Echo Practice")
}
