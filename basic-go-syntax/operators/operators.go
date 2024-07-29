package operators

func IncrementIntByAValue(i *int, o int) {
	*i = *i + o
}

func IncrementIntByAValueShorthand(i *int, o int) {
	*i += o // Shorthand allows supported mathematical operators, such as multiplication (*) and division (/)
}

func IncrementIntByOne(i *int) {
	*i++
}

func DecrementIntByOne(i *int) {
	*i--
}

/*
Copy this function to ./cmd/exec/main.go file to see the results
func main() {
	var i int = 5

	operators.IncrementIntByAValue(&i, 10)
	operators.IncrementIntByOne(&i)
	operators.IncrementIntByAValueShorthand(&i, 8)
	operators.DecrementIntByOne(&i)

	fmt.Println("Look, i was 5 and now it is", i)
}
*/
