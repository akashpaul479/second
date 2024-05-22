package loops

import "fmt"

/*the below code will print square box
* * * * *
* * * * *
* * * * *
* * * * *
* * * * *
 */
func Loops4() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i > 0 && i < 5-1 && j > 0 && j > 5-1 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("* ")
			}
		}
		fmt.Println()
	}
}
