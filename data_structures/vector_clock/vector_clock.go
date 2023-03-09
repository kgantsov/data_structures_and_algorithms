package vectorclock

type VectorClock struct {
	clock []int64
}

func NewVectorClock(number_of_nodes int) *VectorClock {
	clock := make([]int64, number_of_nodes)

	return &VectorClock{
		clock: clock,
	}
}

func (vc *VectorClock) GetClock() []int64 {
	clock := make([]int64, len(vc.clock))

	for i, v := range vc.clock {
		clock[i] = v
	}

	return clock
}

func (vc *VectorClock) Increment(node_id int) {
	vc.clock[node_id]++
}

func (vc *VectorClock) Update(clock []int64) {
	for i := 0; i < len(vc.clock); i++ {
		if clock[i] >= vc.clock[i] {
			vc.clock[i] = clock[i]
		}
	}
}

func (vc *VectorClock) Compare(clock []int64) int {
	equal := true
	smallerOrEqual := true

	for i := 0; i < len(vc.clock); i++ {
		if vc.clock[i] != clock[i] {
			equal = false
		}
		if vc.clock[i] > clock[i] {
			smallerOrEqual = false
		}
	}

	if equal {
		return 0
	}

	if smallerOrEqual && !equal {
		return -1
	}

	return 1
}
