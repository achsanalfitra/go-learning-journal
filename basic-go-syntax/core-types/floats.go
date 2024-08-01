package coretypes

import "fmt"

func ComplareFloatAccuracy(i int) {
	var resultF32 float32 = float32(i) / 999 * 999
	var resultF64 float64 = float64(i) / 999 * 999
	fmt.Println("Your prompt is", i)
	fmt.Printf("The result in float 32 is %f and the result in float64 is %f", resultF32, resultF64)
}
