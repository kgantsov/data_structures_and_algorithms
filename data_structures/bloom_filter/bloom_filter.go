package bloom_filter

import (
	"math"
)

type BloomFilter struct {
	Capacity    int
	Probability float64
	size        int
	hashTable   []int
	hashesNum   int
}

func NewBloomFilter(probability float64, capacity int) *BloomFilter {
	bloomFilter := new(BloomFilter)
	bloomFilter.Probability = probability
	bloomFilter.Capacity = capacity

	bloomFilter.size = int(math.Ceil(
		(float64(capacity) * math.Log(probability)) / math.Log(1/math.Pow(2, math.Log(2))),
	))

	bloomFilter.hashTable = make([]int, bloomFilter.size)
	bloomFilter.hashesNum = int(math.Round(math.Log(2) * float64(bloomFilter.size/capacity)))

	return bloomFilter
}

func (b *BloomFilter) hashFunc(key string, seed int) int {
	h := 0
	for _, char := range key {
		h += int(char) + seed
	}
	return h % b.size
}

func (b *BloomFilter) Add(key string) {
	for seed := 1; seed <= b.hashesNum+1; seed++ {
		b.hashTable[b.hashFunc(key, seed)] = 1
	}
}

func (b *BloomFilter) Check(key string) bool {
	for seed := 1; seed <= b.hashesNum+1; seed++ {
		if b.hashTable[b.hashFunc(key, seed)] == 0 {
			return false
		}
	}

	return true
}

func (b *BloomFilter) GetSize() int {
	return b.size
}
