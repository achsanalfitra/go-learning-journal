package coretypes

import "fmt"

func CompareStringTypes() {
	var rawString string = `Totally rad \n`
	var interpretedString string = "Totally rad \n"

	fmt.Print("This is rawString:", rawString)
	fmt.Println()
	fmt.Print("This is interpretedString:", interpretedString)
}
