package main

import (
	"strconv"
	"fmt"
)

func main() {


	var array [10]string
	for i:=0; i<len(array); i++ {
		array[i] = strconv.Itoa(i)
	}

	var slice []string
	for i:=0; i<10; i++ {
		slice = append(slice, strconv.Itoa(i))
	}

	myMap := make(map[int]string)
	for i:=0; i<10; i++ {
		myMap[i] = strconv.Itoa(i)
	}


	for i, v := range array{
		fmt.Printf("Array i %v - elem %v\n", i, v)
	}

	fmt.Println("---")

	for i, v := range slice{
		fmt.Printf("Slice i %v - elem %v\n", i, v)
	}
	fmt.Println("---")

	for i, v := range myMap{
		fmt.Printf("Map i %v - elem %v\n", i, v)
	}
}