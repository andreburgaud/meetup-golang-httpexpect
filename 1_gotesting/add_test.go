package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tt := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Add2and3", 2, 3, 5},
		{"Add0and3", 0, 3, 3},
		{"Add-100and2", -100, 2, -98},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) { // subtest
			res := Add(tc.a, tc.b)
			if res != tc.expected {
				t.Fatalf("%s should give %d; got %d", tc.name, tc.expected, res)
			}
		})
	}

}
