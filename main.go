package main

import "fmt"

func IsPrime(n int, prime []int) bool {
	if n < 4 {
		return true
	}

	for i := range prime {
		if n%prime[i] == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("start!!")
	n := 100
	prime := []int{}

	for i := 2; i < n; i += 1 {
		if IsPrime(i, prime) {
			prime = append(prime, i)
			fmt.Printf("%d,", i)
		}
	}

	fmt.Println("\nfinish!!")
}
