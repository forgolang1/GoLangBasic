package main

import "fmt"

func main(){

	//result := func(x, y int)(int){
	//	return x+y
	//}(10,20)
	//
	//fmt.Println(result)
	//fmt.Println(sum(1,3))

	//closure("create a new text")

	closure()

	oneMore := func() int{
		z := closure() + 5
		return z
	}

	fmt.Println(oneMore())


	//На 18 минуте 12 видеоурока - ошибка
	//печать значения 100 выполняется не в процессе выполнения кода
	//а за счет того что функция closure вызывается два раза

}

func closure() int {

	t := 100

	result := func() {
		fmt.Println(t)
	}

	//t := 10

	second := func() {
		fmt.Println(t)
	}

	result()
	second()

	return t

}

//func closure(s string) {
//
//	text := "sentence " + " - " + s
//
//	fmt.Println(text)
//
//	result := func() string{
//		fmt.Println(text + "some another text")
//		return text + "some another text"
//	}
//
//	s = "_abc"
//	sq := func() {
//		fmt.Println("ahahah" + s)
//	}
//
//	sq()
//	fmt.Println(result())
//
//}

//func sum(a, b int) int {
//
//	func(s string){
//		fmt.Println(s)
//	}("some text")
//
//	return a + b
//}

