package main

import "fmt"

//...params

func main() {

	sliceOfInt := [][]int{{34,35,38}, {1,2,1}}

	value := calculate(sliceOfInt...)
	fmt.Println(value)
}

func calculate(numbers ...[]int) int {
	result := 0

	//fmt.Printf("%T", numbers)

	for _, v := range numbers {
		for _, v2 := range v {
			result += v2
		}
	}

	return result
}