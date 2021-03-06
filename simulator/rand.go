package simulator

import (
	"math/rand"
	"time"
)

const (
	Partitions = 1000
)

// RandSplitVolume get random partitions of a number
// TODO: there might be better way to get partitions through a combinatorial algo?
// this works for  fake volume
func RandSplitVolume(amount int64) []int64 {
	a := amount
	minPartition := int(amount / Partitions)
	rand.Seed(time.Now().UnixNano())
	amounts := make([]int64, 0, Partitions)
	for a > 0 {
		max := int(a)
		if max > minPartition {
			max = minPartition
		}
		n := int64(rand.Intn(max) + 1)
		amounts = append(amounts, n)
		a = a - n
	}
	return amounts
}
