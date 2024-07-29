package controls

import "fmt"

func PleaseJustInsertFive(i int) {
	if i == 5 {
		fmt.Println("Thank you!")
	}

	fmt.Println("NO! You did not put five, why?")
}

func DoesYourIntEqualToFive(i int) {
	if i == 5 {
		fmt.Println("Yes! Your int equals five")
	} else {
		fmt.Println("Nope")
	}
}

func EitherFiveOrTen(i int) {
	if i == 5 {
		fmt.Println("Your int is five")
	} else if i == 10 {
		fmt.Println("Your int is ten")
	} else {
		fmt.Println("I'm sorry, your int is neither")
	}
}
