package main

import (
	"fmt"

	"github.com/srisubhash011/bookmanagement/services"
)

func main() {
	fmt.Println("Welcome to Number Services")

	var primeNumbers = services.GetPrimeNumber(10)
	fmt.Println("Prime Numbers:", primeNumbers)

	var fibonacci = services.GetFibonacci(10)
	fmt.Println("Fibonacci Sequence:", fibonacci)

	var oddNumbers = services.GetOddNumber(10)
	fmt.Println("Odd Numbers:", oddNumbers)
}
