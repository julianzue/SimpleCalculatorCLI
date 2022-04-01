package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func main() {

	var num1 int
	var num2 int
	var result int
	var op string

	fmt.Print("Type number 1: ")
	fmt.Scan(&num1)

	fmt.Print("Type Operator: ")
	fmt.Scan(&op)

	fmt.Print("Type number 2: ")
	fmt.Scan(&num2)

	if op == "+" {
		result = num1 + num2
	} else if op == "-" {
		result = num1 - num2
	} else if op == "/" {
		result = num1 / num2
	} else if op == "*" {
		result = num1 * num2
	}

	num1str := strconv.Itoa(num1)
	num2str := strconv.Itoa(num2)
	resstr := strconv.Itoa(result)

	fmt.Println("")
	color.Cyan(num1str + " " + op + " " + num2str + " = " + resstr)

}
