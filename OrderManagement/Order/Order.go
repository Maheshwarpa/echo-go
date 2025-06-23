package Order

import "fmt"

type Order struct {
	OrderId         string `json:"OrderId"`
	ProductName     string `json:"ProductName"`
	ProductCategory string `json:"ProductCategory"`
	Quantity        int    `json:"Quantity"`
}

var OrderList map[string]Order

func InitalizeOrderList() {
	OrderList = make(map[string]Order)
}

func AddOrder(order Order) error {
	if _, exists := OrderList[order.OrderId]; exists {
		return fmt.Errorf("Already Present, Please change OrderId")
	} else {
		OrderList[order.OrderId] = order
		return nil
	}
}
