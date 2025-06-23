package Order

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteOrder(c echo.Context) error {

	id := c.Param("id")
	fmt.Println(id)
	delete(OrderList, id)
	if _, exists := OrderList[id]; exists {
		c.JSON(http.StatusFound, map[string]string{
			"error": "Unable to delete the record",
		})

		return fmt.Errorf("Unable to delete")
		//return
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"message": "Record has been Deleted",
		})
		return nil
	}

	//return nil
}
