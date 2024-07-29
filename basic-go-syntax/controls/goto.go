package controls

import "fmt"

func GOTOSTOPWhen100(i int) {
	for {
		if i < 100 {
			i++
		} else if i > 100 {
			i--
		} else {
			fmt.Println("STOP!")
			goto STOP
		}
	}
STOP:
	fmt.Println("You've reached your destination")
}
