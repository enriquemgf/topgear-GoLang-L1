package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)
const error string = "Not a number"
var conversionDictionary = map[rune]string{
		'1':"one",
		'2':"two",
		'3':"three",
		'4':"four",
		'5':"five",
		'6':"six",
		'7':"seven",
		'8':"eight",
		'9':"nine",
		'0':"zero",
		}
		
func num2words(number string)string{
	var literalNumber string = ""
	for _, c := range number {
		literalNumber = literalNumber + " " + conversionDictionary[c]
	}
	return strings.Trim(literalNumber, " ")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(error)
		return
	}
	for _,numberAsString := range args{
		var digitCheck = regexp.MustCompile(`^[0-9]+$`)
		if !digitCheck.MatchString(numberAsString) {
			fmt.Println(error)
			continue
		}
		fmt.Print(num2words(numberAsString))
	}
	
}