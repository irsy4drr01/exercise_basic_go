package modules

import "fmt"

type User struct {
	Name string
	Age  int
}

func (user User) Get() {
	fmt.Println(user)
}

func (user User) Create() {
	fmt.Println("create data user")
}
