package main

import "fmt"

func main() {

	x := 100
	fmt.Println(x)

	someFunc()
	someFuncTwo()

	fmt.Println(miniGlobal)

	dialog()
}

func someFunc() {
	x := 150
	fmt.Println(x)
	fmt.Println(miniGlobal)
}

func someFuncTwo() {

	//fmt.Println(point)
	var test = "ok!"
	var point = 4.4

	fmt.Println(test)
	fmt.Println(point)

}

var miniGlobal = true

func dialog() {
	var x = "Hi"
	var y = "Hello"
	fmt.Println(x)
	{
		fmt.Println(y)
		Z := "How are you?"
		fmt.Println(Z)
	}
	//fmt.Println(Z)
}
