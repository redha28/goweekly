package task1

import "fmt"

func PrintTriangle(rows int) {
	for i := range rows {
		for range i {
			fmt.Print(" ")
		}
		// fmt.Print(i)
		stars := 2*(rows-i) - 1
		for range stars {
			fmt.Print("*")
		}

		fmt.Printf("\n")
	}
}
