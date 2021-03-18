package main

import (
	"fmt"
	"os"
	"strings"
)

var conversionDictionary = map[string]string{
	"one":"1",
	"two":"2",
	"three":"3",
	"four":"4",
	"five":"5",
	"six":"6",
	"seven":"7",
	"eight":"8",
	"nine":"9",
	"zero":"0",
}

func words2num(source string) string {
	var result string = ""
	words:= strings.Fields(source)
	for _,num := range words {
		if val, ok := conversionDictionary[num]; ok {
			result += val
		} else {
			return "Invalid input"
		}
	}
	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error in number of args")
		return 
	}
	arg := args[0]
	result := words2num(strings.ToLower(arg))
	fmt.Println(result)
}