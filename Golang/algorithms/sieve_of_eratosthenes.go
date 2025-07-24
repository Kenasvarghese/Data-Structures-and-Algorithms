package algorithms

// Sieve of Eratosthenes is an algorithm to find all prime numbers up to a given number n.
// It works by iteratively marking the multiples of each prime number as non-prime.
// The algorithm starts with the first prime number, 2, and marks all its multiples as non-prime.
// It then moves to the next unmarked number, 3, and marks all its multiples as non-prime.
// This process continues until all numbers up to n have been processed.
// The remaining unmarked numbers are the prime numbers.
func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)

	// initialize all numbers as prime
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	// mark non-prime numbers as false
	for i := 2; i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}

	// collect all prime numbers
	result := make([]int, 0)
	for i := 2; i <= n; i++ {
		if primes[i] {
			result = append(result, i)
		}
	}
	return result
}