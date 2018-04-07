package main

import (
	"fmt"
)

func main() {
LABEL1:
	for i := 0; i < 10000; i++ {
		for {
			fmt.Println(i)
			continue LABEL1 /*goto:  	0...0...0...0...(unlimited)
			  break: 	0
			  contimue:	0 1 2 3 4 5 6 7 8 9 */
		}
	}
}
