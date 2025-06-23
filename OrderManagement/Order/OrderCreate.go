package Order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context) error {
	var order Order
	err := c.Bind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	err1 := AddOrder(order)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err)
		return err1
	}

	c.JSON(http.StatusAccepted, order)
	return nil
}
