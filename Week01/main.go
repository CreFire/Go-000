package main

import "fmt"

type T1 struct {
	String func() string
}

func (T1) Error() string {
	return "T1.Error"
}

type T2 struct {
	Error func() string
}

func (T2) String() string {
	return "T2.String"
}

var t1 = T1{String: func() string { return "T1.String" }}
var t2 = T2{Error: func() string { return "T2.Error" }}

func main() {
	fmt.Println(t1.Error())
	fmt.Println(t1.String())

	fmt.Println(t2.Error())
	fmt.Println(t2.String())

	fmt.Println(t1)
	fmt.Println(t2)
}