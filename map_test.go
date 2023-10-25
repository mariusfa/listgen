package listgen

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
    list := List[int]{1, 2, 3, 4, 5}

    // Map each number to its square
    squares := Map(list, func(i int) int {
        return i * i
    })

    expected := List[int]{1, 4, 9, 16, 25}
    if !reflect.DeepEqual(squares, expected) {
        t.Fatalf("Map function failed. Expected %v, got %v", expected, squares)
    }
}

func TestEmptyMap(t *testing.T) {
	list := List[int]{} // empty list

	// Map each number to its square
	squares := Map(list, func(i int) int {
		return i * i
	})

	if len(squares) != 0 {
		t.Fatalf("Map function failed. Expected empty list, got %v", squares)
	}
}

