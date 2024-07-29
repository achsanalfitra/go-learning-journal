package controls

import (
	"errors"
	"fmt"
)

func LoopingWithThreeStatements(initial int, conditional int, final int) {
	var err error

	if initial <= conditional && final < 0 {
		err = errors.New("error: you cannot do that")
	} else if initial >= conditional && final > 0 {
		err = errors.New("error: please do not do that")
	} else if conditional == 0 {
		err = errors.New("error: please do not use 0 as conditional")
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := initial; i <= conditional; i += i {
		fmt.Println(i)
	}
}

func BreaksWhenIntReachesAHundred(i int) {
	for {
		if i >= 100 {
			break
		}

		fmt.Println(i)
		i++
	}
}

func RangeOverMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func RangeOverMapWithoutKey(m map[string]string) {
	for _, value := range m {
		fmt.Println(value)
	}
}
