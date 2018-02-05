// https://projecteuler.net/problem=1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

package euler

// sum of arithmetic series given:
// first element (a1)
// number of elements (n)
// series difference (d)
func aritmetic_sereies_sum(a1, n, d int) int {
    // an = a1 + (n - 1)*d
    // Sn = n*(a1 + an)/2
    return n*a1 + n*(n - 1)*d/2
}

// sum of multiples of m smaller than n
func sum_of_multiples_under_n(m, n int) int {
    // the sum of multiples of m under n is the sum of the arithmetic
    // series that star with m, has difference of m, and has n/m elements
    return aritmetic_sereies_sum(m, n/m, m)
}

// greatest common divisor
func gcd(x, y int) int {
    if x < y {
        var tmp = y
        y = x
        x = tmp
    }

    for y > 0 {
        var tmp = x%y
        x = y
        y = tmp
    }
    return x
}

// least common multiplier 
func lcm(x, y int) int {
    return x*y/gcd(x,y)
}

func Problem1() int {
    // summaries the multiples of 3 and 5 under 1000, and removes the ones counted twice
    return sum_of_multiples_under_n(3, 1000) + sum_of_multiples_under_n(5, 1000) - sum_of_multiples_under_n(lcm(3,5), 1000)
}

