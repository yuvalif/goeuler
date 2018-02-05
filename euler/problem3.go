// https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package euler

import "math"

// return all prime factors in a vector, ordered from smallest to largets
func prime_factors(x uint64) (primes []uint64) {
    upper_bound := uint64(math.Sqrt(float64(x)))

    // factors of x are in the range [2,sqrt(x)]
    for i := uint64(2); i <= upper_bound; i++ {
        if x%i == 0 {
            // i is a factor of x
            is_prime := true
            for _, p := range primes {
                // i is a prime factor of x if none of the already
                // calculated prime factors of x is a factor of i
                if i%p == 0 {
                    is_prime = false
                    break
                }
            }
            if is_prime {
                primes = append(primes, i)
            }
        }
    }

    return
}

func Problem3() uint64 {
    primes := prime_factors(600851475143)
    if len(primes) > 0 {
        return primes[len(primes)-1]
    }
    return 0
}

