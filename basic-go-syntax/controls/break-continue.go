package controls

import "fmt"

func SkipsWhenNLessThan100(i int) {
	for {
		if i < 100 {
			i++
			fmt.Println("Skipped")
			continue
		} else if i > 100 {
			fmt.Println("You are overreaching! Your int is now", i)
			i--
		}

		if i == 100 {
			fmt.Println("You got me!")
			break
		} else {
			continue
		}
	}
}
