//A naive implementation of the partition function
package partition

import (
	"sort"
)

type Sum []int

func newSum(n int) Sum {
	return Sum([]int {n})
}

func (s Sum) Append(n int) Sum {
	result := append(s, n)
	sort.Ints(result)
	return result
}

type Partition []Sum

func newPartition(s Sum) Partition {
	return Partition([]Sum {s})
}

func (p Partition) Append(n int) Partition {
	var result Partition
	for i := range p {
		result = append(result, p[i].Append(n))
	}
	return result
}

func (p Partition) NotStartingWithLess(n int) Partition {
	result := Partition(make([]Sum, 0))
	for i := range p {
		if p[i][0] > n {
			result = append(result, p[i])
		}
	}
	return result
}

func PartitionNaive(n int) int {
	return len(partitions(n))
}

func partitions(n int) Partition {
	if n == 0 {
		return newPartition(newSum(0))
	}

	if n == 1 {
		return newPartition(newSum(1))
	}

	result := Partition(make([]Sum, 0))

	for i := 1; i < n; i++ {
		result = append(result, partitions(n-i).NotStartingWithLess(i-1).Append(i)...)
	}

	return append(result, newPartition(newSum(n))...)
}
