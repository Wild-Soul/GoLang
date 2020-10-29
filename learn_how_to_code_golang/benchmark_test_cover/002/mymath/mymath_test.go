package mymath

import "testing"
import "fmt"

func TestCenteredAvg(t *testing.T) {
	inputs := [][]int{{1, 2, 3, 4, 111}, {-111, 2, 3, 4, 222}, {1, 2}, {}}
	outputs := []float64{3, 3, 0, 0}

	for i, te := range inputs {
		fmt.Println(te)
		if CenteredAvg(te) != outputs[i] {
			t.Error("Expected:", outputs[i], "--- Got:", CenteredAvg(te))
		}
	}
}
