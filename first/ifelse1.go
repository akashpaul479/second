package first

import "fmt"

func First2() {
	var a int = 10
	if a%2 == 0 {
		fmt.Printf("a is even")
	} else {
		fmt.Printf("a is odd")
	}
}
