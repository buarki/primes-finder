package main

// findPrimesUpTo implements the Sieve of Eratosthenes
func findPrimesUpTo(limit int) []int {
	primes := make([]bool, limit+1)
	for i := range primes {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	for i := 2; i*i <= limit; i++ {
		if primes[i] {
			for j := i * i; j <= limit; j += i {
				primes[j] = false
			}
		}
	}
	var foundPrimes []int
	for number, isPrime := range primes {
		if isPrime {
			foundPrimes = append(foundPrimes, number)
		}
	}
	return foundPrimes
}
