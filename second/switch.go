package second

import "fmt"

func Second() {
	fmt.Printf("Enter a number:")
	var a, b int
	fmt.Scanln(&a, &b)
	switch a {
	case 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20:
		fmt.Print("the value is even")
	case 1, 3, 5, 7, 9, 11, 13, 15, 17, 19:
		fmt.Print("the value is odd")
	default:
		fmt.Print("the value is undefined")

	}
}
