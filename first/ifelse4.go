package first

import "fmt"

func First4() {
	fmt.Printf("Enter number :")
	var a int
	fmt.Scanln(&a)
	if a < 0 || a > 100 {
		fmt.Print("Enter valid number")
	} else if a >= 0 && a < 50 {
		fmt.Print("FAIL")
	} else if a > 50 && a < 60 {
		fmt.Print("D Grade")
	} else if a > 60 && a < 70 {
		fmt.Print("C Grade")
	} else if a > 70 && a < 80 {
		fmt.Print("B Grade")
	} else if a > 80 && a < 90 {
		fmt.Print("A Grade")
	} else if a > 90 && a < 100 {
		fmt.Print("O Grade")
	}
}
