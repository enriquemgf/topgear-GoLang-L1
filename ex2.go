package main

import (
	"fmt"
	"strconv"
	"os"
	"math"
)
func maxPower(M float64, N int) int{
	k := 0
	for {
		if M <= math.Pow(float64(N), float64(k)) {
			break
		}
		k++
	}
	return k
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Wrong number of arguments")
		return
	}
	M, errM := strconv.Atoi(args[0])
	if errM != nil {
		fmt.Println("Invalid number")
		return
	}
	N, errN := strconv.Atoi(args[1])
	if errN != nil {
		fmt.Println("Invalid number")
		return
	}
	if M < 1 || N < 1{
		fmt.Println("Invalid number")
		return 
	}
	
	fmt.Println(maxPower(float64(M),N))
}
