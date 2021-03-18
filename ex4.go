package main

import (
	"fmt"
	"os"
	"strconv"
	"flag"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println("Error in number of args")
		return
	}
	upperFlag := flag.Bool("u", false, "a bool")
	flag.Parse()
	
	decimal := ""
	if *upperFlag {
		decimal = args[1]
	} else {
		decimal = args[0]
	}
	
	if decimalAsInt, err := strconv.Atoi(decimal); err == nil {
		if *upperFlag {
			fmt.Println(strings.ToUpper(strconv.FormatInt(int64(decimalAsInt), 16)))
		} else {
			fmt.Println(strconv.FormatInt(int64(decimalAsInt), 16))
		}
	} else {
		fmt.Println("Invalid number")
	}
}