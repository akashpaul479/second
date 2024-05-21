package first

import "fmt"

func First3() {
	var a, b int
	fmt.Printf("Enter number a:")
	fmt.Scanln(&a)
	fmt.Print(a)
	fmt.Printf("Enter number b:")
	fmt.Scanln(&b)
	fmt.Print(b)
	if a < b {
		fmt.Printf(" is greater")
	} else {
		fmt.Printf(" is smaller")
	}

}
