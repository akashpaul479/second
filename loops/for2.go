package loops

import "fmt"

/*this below code will print flipped simple pyramid
        *
      * *
    * * *
  * * * *
* * * * *
*/

func Loops2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 2*(5-i)-1; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
