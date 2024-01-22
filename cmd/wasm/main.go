package main

import (
	"syscall/js"
)

func foundPrimes(limit int) []int {
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

func findPrimes() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		limit := args[0].Int()
		x := foundPrimes(limit)

		var res []interface{}
		for _, p := range x {
			res = append(res, p)
		}
		return res
	})
}

func main() {
	js.Global().Set("findPrimes", findPrimes())
	<-make(chan bool)
}
