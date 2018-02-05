// https://projecteuler.net/problem=4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package euler

import "strconv"

func is_palindrome(n uint) bool {

    string_n := strconv.FormatUint(uint64(n), 10)
    left := len(string_n) - 1
    right := 0

    for right < left {
        if string_n[left] != string_n[right] {
            return false
        }
        left--
        right++
    }
    return true
}

func Problem4() (largest_palindrome uint) {
    upper_bound := uint(999)
    lower_bound := uint(100)
    largest_palindrome = 0
    for i := upper_bound; i > lower_bound; i-- {
        for j := i-1; j > lower_bound; j-- {
            if is_palindrome(i*j) && i*j > largest_palindrome {
                largest_palindrome = i*j
            }
        }
    }

    return
}

