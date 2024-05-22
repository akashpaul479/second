package loops

import "fmt"

/*this code will print traingle

      *
    * * *
  * * * * *
* * * * * * *
*/

func Loops3() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 2*(5-i)-1; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
