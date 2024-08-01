package coretypes

import "fmt"

func DisplayInt8(limit int) {
	// int goes negative at maximum length

	var anotherInt int8 = 0
	for i := 0; i <= limit; i++ {
		fmt.Println(anotherInt)
		anotherInt++
	}
}

func DisplayUint8(limit int) {
	// uint starts from zero to maximum length, has max number doubles the size of maximum number of int

	var anotherInt uint8 = 0
	for i := 0; i <= limit; i++ {
		fmt.Println(anotherInt)
		anotherInt++
	}
}
