package main

import (
	"fmt"
)

func main() {

	var num int
	var opt string
	var num2 int
	fmt.Println("enter your number")
	fmt.Scanln(&num)

	for {

		fmt.Println("enter the operation")
		fmt.Scanln(&opt)

		switch {
		case opt == "+":
			num = num + num2
			fmt.Println(num)
		case opt == "-":
			num = num - num2
			fmt.Println(num)
		case opt == "*":
			num = num * num2
			fmt.Println(num)
		default:
			fmt.Println("no operation found")
			return
		}
		fmt.Println("enter your number")
		fmt.Scanln(&num2)

	}

}
