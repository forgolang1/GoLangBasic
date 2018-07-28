package main

import "fmt"

func main() {

	myArray := [5]int{10,20,30,40,50}
	mySlice := []int{100,200,300,400,500}

	fmt.Printf("Array before do something %v\n", myArray)
	DoSomethingArray(myArray)
	fmt.Printf("Array after do something %v\n", myArray)

	fmt.Printf("Slice before do something %v\n", mySlice)
	DoSomethingSlice(mySlice)
	fmt.Printf("Slice after do something %v\n", mySlice)


}

func DoSomethingArray(array [5]int) {
	for i, _ := range array {
		array[i] += 500
	}
	fmt.Printf("Array after finish func %v\n", array)
}

func DoSomethingSlice(slice []int) {
	for i, _ := range slice {
		slice[i] += 10000
	}
	fmt.Printf("Slice after finish func %v\n", slice)
}