//programming Lecture 6
package main

import (
	"fmt"
)

func main() {
	a := []int{5, 8, 7, 3, 4}
	fmt.Println(a)
	BubbleAsort(a)
	BubbleZsort(a)
}

func BubbleAsort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}

func BubbleZsort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	fmt.Println(a)
}
