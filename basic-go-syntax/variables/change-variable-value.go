package variables

func ChangeVariableDirectly(i, o any) any {
	var input any = i
	input = o
	return input
}

func ChangeVariableUsingPointer(i *any, o any) {
	*i = o
}

/*
Copy this function to ./cmd/exec/main.go file to see the results
func main() {
	a := variables.ChangeVariableDirectly(5, 10)

	var b any = 5
	variables.ChangeVariableUsingPointer(&b, 10)

	fmt.Println("Both a and b were 5, but now they are both", a, "and", b)
}
*/
