package main

import "flag"
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
	var n int
	flag.IntVar(&n, "num", 10, "int flag")
	flag.Parse()
	t := time.Now()
	fmt.Printf("%s start!!\n", t)
	prime := []int{}

	for i := 2; i < n; i += 1 {
		if IsPrime(i, prime) {
			prime = append(prime, i)
		}
	}

	t = time.Now()
	fmt.Printf("%s %d番目の素数: %d\n", t, n, prime[len(prime)-1])
}
