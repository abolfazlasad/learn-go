// Question 11.2 — table-driven tests
//
// Write TestAdd with a table of {a, b, want}.
// Use t.Run for each case. Fail with t.Errorf, not panic.
//
// Extra: a helper assertEq that calls t.Helper().
//
// Run:
//   go test 11.2-add.go 11.2-add_test.go

package add

import (
	"fmt"
	"testing"
)

func assertEq(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("a=%d, b=%d", test.a, test.b), func(t *testing.T) {
			assertEq(t, Add(test.a, test.b), test.want)
		})
	}
}
