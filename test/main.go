package main

import (
	"fmt"
)

type CustomerError error

type SecondLevelError CustomerError

type Myint int

type SecondInt Myint

func (m Myint) String() string {
	return "Myint"
}

func (s SecondInt) String() string {
	return "SecondInt"
}

func main() {

	testInt := 10
	testMy := Myint(testInt)
	testS := SecondInt(testMy)
	fmt.Printf("test var %d, %d, type %T\n", testInt, testInt, testInt)

	fmt.Printf("test var %d, %s, type %T\n", testMy, testMy, testMy)

	fmt.Printf("test var %d, %s, type %T\n", testS, testS, testS)
}
