package listgen

import (
    "testing"
)

func TestFilter(t *testing.T) {
    list := List[int]{1, 2, 3, 4, 5}

    evens := list.Filter(func(i int) bool {
        return i%2 == 0
    })

    if len(evens) != 2 || evens[0] != 2 || evens[1] != 4 {
        t.Fatalf("Filter function failed. Expected [2, 4], got %v", evens)
    }
}
