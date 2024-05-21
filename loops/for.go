package loops

import "fmt"

func Loops() {
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			fmt.Print(a, "  ", b, "\n")
		}
	}
}
