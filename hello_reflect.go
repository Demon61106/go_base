package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"11111"b:"22222"`
}

type INT int

type A struct {
	a int
}

type B struct {
	b string
}

type Ita interface {
	String() string
}

func (b B) String() string {
	return b.b
}

func reflectTest1() {
	s := Student{}
	rt := reflect.TypeOf(s)

	fieldName, ok := rt.FieldByName("Name")

	if ok {
		fmt.Println(fieldName.Tag)
	}

	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_Name: ", rt.Name())
	fmt.Println("type_NumField: ", rt.NumField())
	fmt.Println("type_PkgPath :", rt.PkgPath())
	fmt.Println("type_String :", rt.String())

	fmt.Println("type.Kind.String:", rt.Kind().String())
	fmt.Println("type.String()", rt.String())

	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name:=%v \n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice element type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice type.PkgPaht()=", scet.PkgPath())
	fmt.Println("slice type.PkgPaht()=", sct.PkgPath())
}

func reflectTest2() {
	var a INT = 12
	var b int = 14
	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta == tb {
		fmt.Println("ta == tb ")
	} else {
		fmt.Println("ta != tb")
	}

	fmt.Println("ta.Name ", ta.Name())
	fmt.Println("tb.Name ", tb.Name())

	fmt.Println("ta.Kind.String ", ta.Kind().String())
	fmt.Println("tb.Kind.String ", tb.Kind().String())

	s1 := A{1}
	s2 := B{"tata"}
	fmt.Println("reflect.TypeOf(s1).Name() = ", reflect.TypeOf(s1).Name())
	fmt.Println("reflect.TypeOf(s2).Name() = ", reflect.TypeOf(s2).Name())

	fmt.Println("reflect.TypeOf(s1).Kind().String() = ", reflect.TypeOf(s1).Kind().String())
	fmt.Println("reflect.TypeOf(s2).Kind().String() = ", reflect.TypeOf(s2).Kind().String())

	ita := new(Ita)
	var itb Ita = s2
	fmt.Println("reflect.TypeOf(ita).Elem().Name() ", reflect.TypeOf(ita).Elem().Name())
	fmt.Println("reflect.TypeOf(ita).Elem().Kind().String()", reflect.TypeOf(ita).Elem().Kind().String())

	fmt.Println("reflect.TypeOf(itb).Name() ", reflect.TypeOf(itb).Name())
	fmt.Println("reflect.TypeOf(itb).Elem().String()", reflect.TypeOf(itb).Kind().String())

}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) String() {
	fmt.Println("User :", this.Id, this.Name, this.Age)
}

func Info(o interface{}) {
	v := reflect.ValueOf(o)

	t := v.Type()

	fmt.Println("Type: ", t.Name())

	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		switch value := value.(type) {
		case int:
			fmt.Printf(" %6s: %v=%d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf(" %6s: %v=%s\n", field.Name, field.Type, value)
		default:
			fmt.Printf(" %6s: %v=%s\n", field.Name, field.Type, value)
		}
	}
}

func reflectTest3() {
	u := User{1, "Tom", 30}
	Info(u)
}

func reflectTest4() {
	u := User{1, "andes", 20}
	va := reflect.ValueOf(u)
	vb := reflect.ValueOf(&u)

	fmt.Println(va.CanSet(), va.FieldByName("Name").CanSet())
	fmt.Println(vb.CanSet(), vb.Elem().FieldByName("Name").CanSet())

	fmt.Printf("%v\n", vb)
	name := "shine"
	vc := reflect.ValueOf(name)

	vb.Elem().FieldByName("Name").Set(vc)
	fmt.Printf("%v\n", vb)
}

func relfectTest() {
	reflectTest4()
}
