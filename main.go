package main

import "github.com/irsy4drr01/basic_go-1/modules"

func main() {
	// exported dan unexported func
	// sumResult, timesResult := modules.Sum(2, 3)
	// fmt.Println(sumResult, timesResult)

	// defer dan exit
	// defer modules.CloseDbConnection()
	// result, error := modules.QueryGetUser(false)
	// if error != nil {
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println("get data user")

	// pointer
	modules.Pointer()
	modules.NoPointer()

	// var product1 modules.Product
	// product1.Name = "Indomie"
	// product1.Price = 30000
	// product1.Image = "indomie.png"
	// fmt.Println(product1.GetName())

	// products := []modules.Product{
	// 	{
	// 		Name:  "Indomie",
	// 		Price: 30000,
	// 		Image: "indomie.png",
	// 	},
	// 	{
	// 		Name:  "Coca Cola",
	// 		Price: 5000,
	// 		Image: "coca-cola.png",
	// 	},
	// }
	// for _, product := range products {
	// 	fmt.Println(product.Name)
	// }
}
