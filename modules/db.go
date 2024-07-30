package modules

import (
	"errors"
	"fmt"
)

func CloseDbConnection() {
	fmt.Println("close db connection")
}

func QueryGetUser(isError bool) ([]string, error) {
	if isError {
		return []string{}, errors.New("terjadi kesalahan")
	} else {
		return []string{"data1", "data2"}, nil
	}
}
