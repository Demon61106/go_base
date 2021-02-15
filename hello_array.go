package main

import "fmt"

func ArrayTest() {
	a := [3]int{1, 2, 3}
	b := a

	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)

}
