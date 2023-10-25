package listgen

import (
	"testing"
)

func TestList_Add(t *testing.T) {
    l := List[int]{1, 2, 3}
    l.Add(4)
    if len(l) != 4 {
        t.Errorf("Expected length of list to be 4, but got %d", len(l))
    }
    if l[3] != 4 {
        t.Errorf("Expected last element of list to be 4, but got %d", l[3])
    }
}