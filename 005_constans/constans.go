package main

import "fmt"

const applicationName = "App Const"

const (
	_ = iota //0
	_ = iota
	_ = iota
	D = iota + 5 // 3 + 5
	E = iota     // 4
)

func main() {

	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

	const number = 100
	fmt.Println(applicationName)
	fmt.Println(number)

	var numberTwo int = 10
	result := number + numberTwo
	fmt.Println(result)

	x, _, _ := TypesGo()
	fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)

}

func TypesGo() (int, bool, string) {
	return 100, false, "stringType"
}
