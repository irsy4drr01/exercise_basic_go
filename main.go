package main

import (
	"fmt"
	// "github.com/irsy4drr01/basic_go-1/modules"
)

func main() {
	// // exported dan unexported func
	// sumResult, timesResult := modules.Sum(2, 3)
	// fmt.Println(sumResult, timesResult)

	// // defer dan exit
	// defer modules.CloseDbConnection()
	// result, error := modules.QueryGetUser(false)
	// if error != nil {
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println(result)
	// }
	// fmt.Println("get data user")

	// // pointer
	// modules.Pointer()
	// modules.NoPointer()

	// // struct
	// // struct tanpa menggunakan service
	// // menampilkan struct biasa
	// fmt.Println("STRUCT BIASA")
	// var product1 modules.Product
	// product1.Name = "Indomie"
	// product1.Price = 30000
	// product1.Image = "indomie.png"
	// fmt.Println(product1.GetName())
	// fmt.Println()

	// // menampilkan struct dengan objek
	// fmt.Println("STRUCT DENGAN OBJEK")
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
	// fmt.Println()

	// // struct dengan menggunakan service
	// fmt.Println("STRUCT DENGAN MENGGUNAKAN SERVICE")

	// productService := modules.Product{
	// 	Name:  "Basko",
	// 	Price: 15000,
	// 	Image: "bakso.png",
	// }
	// productService.Get()
	// productService.Create()

	// userService := modules.User{
	// 	Name: "Andi",
	// 	Age:  25,
	// }
	// userService.Get()
	// userService.Create()

	// any data type
	// menggunakan interface{} === any
	var response any = 400 // bisa berisi type data apapun
	fmt.Println(response)
}
