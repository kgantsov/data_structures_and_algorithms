package bloom_filter

import (
	"math"
)

type BloomFilter struct {
	capacity    int
	probability float32
	size        int
	hash_table  []int
	hashes_num  int
}

func NewBloomFilter(probability float64, capacity int) *BloomFilter {
	bloomFilter := new(BloomFilter)

	bloomFilter.size = int(math.Ceil(
		(float64(capacity) * math.Log(probability)) / math.Log(1/math.Pow(2, math.Log(2))),
	))

	bloomFilter.hash_table = make([]int, bloomFilter.size)
	bloomFilter.hashes_num = int(math.Round(math.Log(2) * float64(bloomFilter.size/capacity)))

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
	for seed := 1; seed <= b.hashes_num+1; seed++ {
		b.hash_table[b.hashFunc(key, seed)] = 1
	}
}

func (b *BloomFilter) Check(key string) bool {
	for seed := 1; seed <= b.hashes_num+1; seed++ {
		if b.hash_table[b.hashFunc(key, seed)] == 0 {
			return false
		}
	}

	return true
}
