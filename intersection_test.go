package godash

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	a1 := []int{ 1, 2, 3, 4, 5 }
	a2 := []interface{}{ 1, "two", 3, "four" }

	i := Intersection(a1, a2)

	if len(i) != 2 {
		t.Errorf("Expected [1,3], Got %+v", i)
	}
}
