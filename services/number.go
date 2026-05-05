package services

func GetPrimeNumber(n int) []int {
	var primeNumbers []int
	for i := 2; len(primeNumbers) < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}

func GetFibonacci(n int) []int {
	var fibonacci []int
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fibonacci = append(fibonacci, a)
		a, b = b, a+b
	}
	return fibonacci
}

func GetOddNumber(n int) []int {
	var oddNumbers []int
	for i := 1; len(oddNumbers) < n; i += 2 {
		oddNumbers = append(oddNumbers, i)
	}
	return oddNumbers
}
