package loops

import "fmt"

func Loops1() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
