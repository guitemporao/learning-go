package int

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	//You will notice that we're using %d as our format strings rather than %q. That's because we want it
	//to print an integer rather than a string.

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

	sub := Sub(2, 2)
	expectedSub := 0

	if sub != expectedSub {
		t.Errorf("expected '%d' but got '%d'", expectedSub, sub)

	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
