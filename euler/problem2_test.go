// test suite for problem2.go

package euler

import "testing"

var results = map[uint]uint64 {
    0: 0,
	1: 1,
	2: 1,
	3: 2,
	4: 3,
	5: 5,
	6: 8,
    10: 55,
    14: 377,
    20: 6765,
    29: 514229,
    33: 3524578,
    59: 956722026041,
}

func Test_fibonacci(t *testing.T) {
    for arg, expected := range results {
        actual := fibonacci(arg)
        if actual != expected {
            t.Error("Expected", expected, "got", actual)
        }
    }
}

func Test_find_nearest_fibonacci_index(t *testing.T) {
    for expected, arg := range results {
		// no point in checking for "1", since there are two options
		if arg == 1 {
			continue
		}
        actual := find_nearest_fibonacci_index(arg)
        if actual != expected {
            t.Error("Expected", expected, "got", actual)
        }
    }
}

