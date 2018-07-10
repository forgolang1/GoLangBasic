package main

import "fmt"

func main() {

	for x:=1; x<100; x++ {

		if x%10 == 0 {
			fmt.Println("%10", x)
			continue //прерываем текущую итерацию цикла for
			// break - остановка цикла
		}

		fmt.Println(x)
	}
}