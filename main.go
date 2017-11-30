package main

import "fmt"
import "time"

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
	t := time.Now()
	fmt.Printf("%s start!!\n", t)
	n := 1000000
	prime := []int{}

	for i := 2; i < n; i += 1 {
		if IsPrime(i, prime) {
			prime = append(prime, i)
		}
	}

	t = time.Now()
	fmt.Printf("%s %d番目の素数: %d\n", t, n, prime[len(prime)-1])
}
