package main

import "fmt"

func main() {

	var sliceOne []int
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	sliceOne = append(sliceOne, 10)
	sliceOne = append(sliceOne, 20)
	sliceOne = append(sliceOne, 30)
	sliceOne = append(sliceOne, 40)
	sliceOne = append(sliceOne, 50)

	//fmt.Println(len(sliceOne))
	//fmt.Println(cap(sliceOne))

	//fmt.Println(sliceOne[3:])
	//fmt.Println(sliceOne[:3])
	//fmt.Println(sliceOne[2:4])
	//sliceOne[4] = 1000
	//sliceOne[5] = 19

	//fmt.Println(sliceOne)

	slice := make([]int, 5, 19)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	//fmt.Println(slice[5])
	//fmt.Println((sliceOne))

	x := []int{1,2,3}
	fmt.Println(x)
}
