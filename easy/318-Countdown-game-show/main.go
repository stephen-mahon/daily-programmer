package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	target := r.Intn(999)
	smalls := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	larges := []int{25, 50, 75, 100}
	small := flag.Int("s", 6, "total small numbers")
	large := flag.Int("l", 0, "total large numbers")
	flag.Parse()
	if (*small + *large) != 6 {
		log.Fatalf("total numbers must equal 6: smalls (%v) + larges (%v) = %v", *small, *large, *small+*large)
	}
	if *large > 4 {
		log.Fatal("there are only 4 large numbers")
	}
	var nums []int
	count := 0
	for count != *small {
		i := r.Intn(len(smalls))
		nums = append(nums, smalls[i])
		smalls = remove(smalls, i)
		count++
	}
	count = 0
	for count != *large {
		i := r.Intn(len(larges))
		nums = append(nums, larges[i])
		larges = remove(larges, i)
		count++
	}
	fmt.Println("Your numbers are:", nums)
	fmt.Println("The target today is:", target)

	countdownSolver(target, nums)

}

func countdownSolver(target int, nums []int) {
	// Maybe break the target down into its prime factors?
	// Since a prime cannot be found my multiplication we must use another method
	primeFactors := factorPrime(target)
	fmt.Println(primeFactors)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func factorPrime(n int) []int {
	var s []int
	if isPrime(n, 2) {
		s = append(s, 1)
		s = append(s, n)
		return s
	}
	for i := 1; i <= n/2; i++ {
		if n%i == 0 && isPrime(i, 2) {
			s = append(s, i)
		}
	}
	return s
}

func isPrime(p, n int) bool {
	if p == 0 || p == 1 {
		return false
	}

	if p == n {
		return true
	}

	if p%n == 0 {
		return false
	}
	n++
	return isPrime(p, n)
}
