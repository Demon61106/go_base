package main

import "fmt"

func fa(a int) func(i int) int {
	return func(i int) int {
		fmt.Println(&a, a)
		a = a + i
		return a
	}
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		m[stu.Name] = &stus[i]
	}
	fmt.Println(m)
}

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func FuncTest1() {
	var peo People = &Stduent{}
	// peo := Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

func FuncTest() {
	pase_student()
	// f1 := fa(1)
	// fmt.Println(f1(1))
	// fmt.Println(f1(2))

	// f2 := fa(2)
	// fmt.Println(f2(1))
	// fmt.Println(f2(2))
}
