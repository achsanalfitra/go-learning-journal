package main

import (
	"fmt"

	"github.com/achsanalfitra/go-learning-journal/basic-go-syntax/controls"
)

func main() {
	i := []int{2, 1, 3, 4, 6, 5, 7}

	controls.BubbleSortSolution(&i)
	fmt.Println(i)
}
