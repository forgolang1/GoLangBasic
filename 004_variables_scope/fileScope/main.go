package main

import "fmt"

var car = "ford"

func main()  {

	fmt.Println(car)
	checkSpareWheel()

}

func checkSpareWheel() {
	if car == "ford" {
		fmt.Println("Spare wheel exist")
	}
}
