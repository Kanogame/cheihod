package utils

import "fmt"

func UserError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ServerError(err error) {
	if err != nil {
		panic(err)
	}
}
