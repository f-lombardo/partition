package partition

import (
	"testing"
)

func TestPartitionNaive(t *testing.T) {
	executeFor(t, PartitionNaive)
}

func executeFor(t *testing.T, partitionFuncion func(n int) int) {
	cases := []int {1,1,2,3,5,7,11,15,22,30,42,56,77,101,135,176,231,297,385,490,627,792,1002,1255}

	for input, expected := range cases {
		actual := partitionFuncion(input)
		if actual != expected {
			t.Errorf("Partiton(%d) == %d, expected %d", input, actual, expected)
		}
	}
}
