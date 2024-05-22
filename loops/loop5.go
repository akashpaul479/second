package loops

import "fmt"

/* this code will print the * in cross and the other space is cover by numbers
* 0000
1* 111
22* 22
333* 3
4444*
*/

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
