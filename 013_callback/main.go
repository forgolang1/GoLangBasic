package main

import "fmt"

//func main() {
//
//	t := calculate(100)
//
//	//t := func(k int) int {
//	//	return k * 1000
//	//}
//
//	z := t(5)
//	fmt.Println(z)
//
//
//}
//
//func calculate(i int) func(int)int {
//	fmt.Println(i)
//
//	return func(k int) int {
//		return k * 1000
//	}
//}

func calculate(numbers []int, justFunc func(int) int) int {
	result := 0

	for _, n := range numbers {
		fmt.Println(justFunc(n))

		result += justFunc(n)
	}

	return result
}

func main() {
	sOfInt := []int{1, 2, 3, 4}

	x := calculate(sOfInt, func(n int) int {
		return n * 10
	})

	fmt.Println(x)
}
