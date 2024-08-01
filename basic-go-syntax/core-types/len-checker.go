package coretypes

import "fmt"

func RuneLengthChecker(s string) {
	runeS := []rune(s)
	fmt.Println(len(runeS))
}

func DirectLengthChecker(s string) {
	fmt.Println(len(s))
}
