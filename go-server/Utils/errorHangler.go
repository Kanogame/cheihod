package utils

import "fmt"

func UserError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func ServerError(err error) {
	if err != nil {
		panic(err)
	}
}
