// test suite for problem1.go

package euler

import "testing"

func Test_aritmetic_sereies_sum(t *testing.T) {
    a1 := 1
    n := 99
    d := 7
    sum := a1
    an := a1
    for i := 2; i <= n; i++ {
        an += d
        sum += an
    }
    if aritmetic_sereies_sum(a1, n, d) != sum {
        t.Error("Expected ", sum, " got ", aritmetic_sereies_sum(a1, n, d))
    }

    a1 = 71
    n = 999
    d = 13
    sum = a1
    an = a1
    for i := 2; i <= n; i++ {
        an += d
        sum += an
    }
    if aritmetic_sereies_sum(a1, n, d) != sum {
        t.Error("Expected ", sum, " got ", aritmetic_sereies_sum(a1, n, d))
    }
}

