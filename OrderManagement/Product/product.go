package Product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Product struct {
	ProductId       string  `json:"ProductId"`
	ProductName     string  `json:"ProductName"`
	ProductCategory string  `json:"ProductCategory"`
	ProductQuantity int     `json:"ProductQuantity"`
	ProductPrice    float64 `json:"ProductPrice"`
}

var ProductList []Product

func LoadData() {
	file, err := os.Open("data.json")
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Unable to read the data")
		return
	}

	err1 := json.Unmarshal(data, &ProductList)
	if err1 != nil {
		return
	}

	fmt.Println("ProductList is: ", ProductList)
}
