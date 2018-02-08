// https://projecteuler.net/problem=5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// 
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package euler

import "github.com/bxcodec/saint"

type prime_factors_t = map[uint]uint8

// factorize a number into its prime factors
// the returned list is holding the prime factor and its exponent
func factorize(n uint) (v prime_factors_t) {
    v = make(prime_factors_t)
    m := uint(2)
    // 1 is always a factor of n
    v[1] = 1
    for n > 1 {
        if n%m == 0 {
            // m is a factor of n
            _, ok := v[m]
            if (!ok) {
                // insert new prime and set exponent to 1
                v[m] = 1
            } else {
                // prime exists- increment exponent
                v[m]++
            }
            n /= m
        } else {
            // m is not a factor of n
            m++
        }
    }

    return
}

// given a list of prime factors, the original number is calculated
// under the assumptions that each entry in the vector indicate the power of the prime factor
func defactorize(v prime_factors_t) (total uint64) {
    total = 1
    for prime, exponent := range v {
        total *= uint64(saint.Pow(int(prime), int(exponent)))
    }
    return
}

func Problem5() uint64 {
    max_factors := make(prime_factors_t)
    for i := uint(1); i <= 20; i++ {
        // factorize the next number
        v := factorize(i);
        for prime, exponent := range v {
            // check if prime is already in the max factors list
            max_exponent, ok := max_factors[prime]
            if (!ok) {
                // add to list
                max_factors[prime] = exponent
            } else {
                // update to max value between existing and new
                max_factors[prime] = uint8(saint.Max(int(max_exponent), int(exponent)))
            }
        }
    }

    return defactorize(max_factors)
}

