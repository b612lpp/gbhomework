package main

import "fmt"

func main() {

	var percent, deposit float64
	fmt.Println("Input deposit value")
	fmt.Scan(&deposit)
	fmt.Println("Input deposit percent")
	fmt.Scan(&percent)
	for i := 1; i < 6; i++ {
		deposit = deposit + ((deposit * percent) / 100)
		fmt.Println("after ", i, "year deposit gonna be ", deposit)
	}

}
