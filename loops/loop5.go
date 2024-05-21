package loops

import "fmt"

func Loop5() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				fmt.Printf("* ")
			} else {
				fmt.Printf("%d", i)
			}
		}
		fmt.Println()
	}
}
