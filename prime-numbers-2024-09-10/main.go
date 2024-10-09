package main

import (
	"fmt"
)

/*
Write a Go program that prints the numbers from 1 to 50. 
But for multiples of 3, print "Fizz" instead of the number, 
and for the multiples of 5, print "Buzz". For numbers which are multiples of both 3 and 5, print "FizzBuzz".
Additionally, if the number is prime, print "Prime" instead of any other output (i.e., it overrides "Fizz" and "Buzz").
*/
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i :=2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}


func FizzBuzzChecker(numbers []int) {
	for _, number := range numbers{
		switch {
		case IsPrime(number):
			fmt.Println("Prime")
		case number %5 == 0 && number %3 == 0:
			fmt.Println("FizzBuzz")
		case number %5 == 0:
			fmt.Println("Buzz")
		case number %3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(number)
		}
	
	}
}



func main() {

	var numbers []int
	for i:=1; i < 51; i++ {
		numbers = append(numbers, i)
	}
	FizzBuzzChecker(numbers)
}