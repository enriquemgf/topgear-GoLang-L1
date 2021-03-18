package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)
	if len(args) != 3 {
		fmt.Println("Error in number of args")
		return 
	}
	complex1,err1 := strconv.ParseComplex(args[0], 64)
	complex2,err2 := strconv.ParseComplex(args[2], 64)
	bynaryop := args[1]
	
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid number")
		return
	}

	switch bynaryop {
	case "+":
		fmt.Println(complex1 + complex2)
	case "-":
		fmt.Println(complex1 - complex2)
	case "*":
		fmt.Println(complex1 * complex2)
	case "/":
		fmt.Println(complex1 / complex2)
	}

}