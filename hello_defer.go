package main

import "fmt"

func f1() (r int) {
	defer func() {
		r += 5
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t += 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r += 5
	}(r)
	return 1
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func DeferTest1() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func DeferTest() {
	fmt.Println("test f1 ", f1())
	fmt.Println("test f2 ", f2())
	fmt.Println("test f3 ", f3())

	DeferTest1()
}
