package dog

import (
	"testing"
)

type test struct {
	x int
	y int
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}

func TestYears(t *testing.T) {
	tests := []test{{1, 7}, {2, 14}, {3, 21}}

	for _, te := range tests {
		if Years(te.x) != te.y {
			t.Error("Expected", te.y, " Got:", Years(te.x))
		}
	}
}

func TestYearsTwo(t *testing.T) {
	tests := []test{{1, 7}, {2, 14}, {3, 21}}

	for _, te := range tests {
		if YearsTwo(te.x) != te.y {
			t.Error("Expected", te.y, " Got:", YearsTwo(te.x))
		}
	}
}
