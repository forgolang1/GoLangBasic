package main

import "fmt"

func main() {

	defer theFour()
	defer theThree()
	defer theTwo()
	 theOne()

}

func theOne() {
	fmt.Println("One")
}

func theTwo() {
	fmt.Println("Two")
}

func theThree() {
	fmt.Println("Three")
}

func theFour() {
	fmt.Println("Four")
}

