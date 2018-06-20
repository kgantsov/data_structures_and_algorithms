package heap

type Heap struct {
	nodes []int
	size  int
}

func NewHeap(args ...int) *Heap {
	heap := new(Heap)
	heap.nodes = make([]int, 1)
	heap.size = 0

	for _, v := range args {
		heap.Push(v)
	}

	return heap
}

func (h *Heap) moveUp(size int) {
	for size/2 > 0 {
		if h.nodes[size] > h.nodes[size/2] {
			h.nodes[size], h.nodes[size/2] = h.nodes[size/2], h.nodes[size]
		}

		size /= 2
	}
}

func (h *Heap) maxChild(index int) int {
	if index*2+1 > h.size {
		return index * 2
	} else {
		if h.nodes[index*2] > h.nodes[index*2+1] {
			return index * 2
		}

		return index*2 + 1
	}
}

func (h *Heap) moveDown(index int) {
	for index*2 <= h.size {
		maxChild := h.maxChild(index)

		if h.nodes[index] < h.nodes[maxChild] {
			h.nodes[index], h.nodes[maxChild] = h.nodes[maxChild], h.nodes[index]
		}

		index = maxChild
	}

}

func (h *Heap) Push(value int) {
	h.nodes = append(h.nodes, value)
	h.size++
	h.moveUp(h.size)

}

func (h *Heap) Pop() int {
	value := h.nodes[1]
	h.nodes[1] = h.nodes[h.size]
	// h.nodes[h.size] = 0
	h.nodes = h.nodes[:h.size]
	h.size--
	h.moveDown(1)
	return value
}
