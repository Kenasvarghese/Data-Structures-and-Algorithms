def sieve_of_eratosthenes(n: int) -> list[int]:
    """
    Sieve of Eratosthenes is an algorithm to find all prime numbers up to a given number n.
    It works by iteratively marking the multiples of each prime number as non-prime.
    The algorithm starts with the first prime number, 2, and marks all its multiples as non-prime.
    It then moves to the next unmarked number, 3, and marks all its multiples as non-prime.
    This process continues until all numbers up to n have been processed.
    The remaining unmarked numbers are the prime numbers.
    """

    # initialize all numbers as prime
    primes = [True] * (n+1)
    primes[0] = False
    primes[1] = False

    # mark non-prime numbers as false
    for i in range(2,n+1):
        if primes[i]:
            for j in range(i*i, n+1, i):
                primes[j] = False

    # collect all prime numbers
    return [i for i in range(2, n+1) if primes[i]]


if __name__ == "__main__":
    print(sieve_of_eratosthenes(100))