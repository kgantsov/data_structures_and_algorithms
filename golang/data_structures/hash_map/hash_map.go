package hash_map

import (
	"hash/fnv"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func remove(s []Node, i int) []Node {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

type Node struct {
	key   string
	value interface{}
}

type HashMap struct {
	capacity  int
	length    int
	hashTable [][]Node
}

func NewHashMap() *HashMap {
	m := new(HashMap)
	m.capacity = 2

	m.hashTable = make([][]Node, m.capacity)
	for i := range m.hashTable {
		m.hashTable[i] = make([]Node, 0)
	}

	return m
}

func (m *HashMap) Len() int {
	return m.length
}

func (m *HashMap) resize(capacity int) {

	// Create a new hashTable for values to be moved
	hashTable := make([][]Node, capacity)
	for i := range hashTable {
		hashTable[i] = make([]Node, 0)
	}

	// Copy all keys and values from the old hashTable to a new one
	for i := range m.hashTable {
		for j := range m.hashTable[i] {
			node := &m.hashTable[i][j]

			index := int(hash(node.key)) % capacity

			found := false
			for i := range hashTable[index] {
				node := &hashTable[index][i]
				if node.key == node.key {
					node.value = node.value
					found = true
				}
			}

			if !found {
				hashTable[index] = append(hashTable[index], Node{key: node.key, value: node.value})
			}
		}
	}

	m.hashTable = hashTable
	m.capacity = capacity
}

func (m *HashMap) set(key string, value interface{}) bool {
	index := int(hash(key)) % m.capacity

	// Make hash table two times bigger if it more than half full to have less change having collisions
	if m.load() > 0.5 {
		m.resize(m.capacity * 2)
	}

	for i := range m.hashTable[index] {
		node := &m.hashTable[index][i]
		if node.key == key {
			node.value = value
			return true
		}
	}

	node := Node{key: key, value: value}
	m.hashTable[index] = append(m.hashTable[index], node)
	m.length++

	return true
}

func (m *HashMap) get(key string) (interface{}, bool) {
	index := int(hash(key)) % m.capacity

	for i := range m.hashTable[index] {
		node := &m.hashTable[index][i]
		if node.key == key {
			return node.value, true
		}
	}

	return nil, false
}

func (m *HashMap) del(key string) bool {
	index := int(hash(key)) % m.capacity

	// Make hash table two times smaller if it less than quarter full to release memory
	if m.load() < 0.25 {
		m.resize(m.capacity / 2)
	}

	for i := range m.hashTable[index] {
		node := &m.hashTable[index][i]
		if node.key == key {
			m.hashTable[index] = remove(m.hashTable[index], i)
			m.length--
			return true
		}
	}

	return false
}

func (m *HashMap) load() float64 {
	return float64(m.length) / float64(m.capacity)
}
