package Order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateOrder(c echo.Context) error {

	var order Order
	err := c.Bind(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	OrderList[order.OrderId] = Order{
		OrderId: order.OrderId, ProductName: order.ProductName, ProductCategory: order.ProductCategory, Quantity: order.Quantity,
	}

	c.JSON(http.StatusAccepted, OrderList[order.OrderId])

	return nil
}
