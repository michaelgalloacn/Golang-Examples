package main

import (
	"fmt"

	errorhandling "github.com/michaelgalloacn/Golang-Examples/errorhandling"
)

func main() {
	str, err := errorhandling.StringWithError(true)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(str)

}
