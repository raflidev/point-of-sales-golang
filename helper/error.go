package helper

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func DD(data interface{}) {
	fmt.Println(data)
}
