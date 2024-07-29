package controls

import "fmt"

func EitherFiveOrTenUsingSwitches(i int) {
	switch i {
	case 5:
		fmt.Println("Your int is five")
	case 10:
		fmt.Println("Your int is ten")
	default:
		fmt.Println("I'm sorry, your int is neither")
	}
}

func IsYourStringPartOfABC(i string) {
	switch i {
	case "A", "B", "C":
		fmt.Println("Yes it does!")
	default:
		fmt.Println("No-uh!")
	}
}
