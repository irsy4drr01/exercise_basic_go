package modules

import "fmt"

type Product struct {
	Name  string
	Price int
	Image string
}

func (product Product) Get() {
	fmt.Println(product)
}

func (product Product) Create() {
	fmt.Println("create data product")
}

func (product Product) GetName() Product {
	return product
}
