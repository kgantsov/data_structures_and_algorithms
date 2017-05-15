package insert

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	list := []int{6, 8, 45, 1, 84, 149, 9, 17}
	expectedList := []int{1, 6, 8, 9, 17, 45, 84, 149}

	Sort(list)

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}

	list = []int{50, 66, 20, 22, 16, 11, 4, 91, 51, 1, 68, 87, 10, 54, 81, 39, 15, 49, 68, 74, 54, 27, 32, 23}
	expectedList = []int{1, 4, 10, 11, 15, 16, 20, 22, 23, 27, 32, 39, 49, 50, 51, 54, 54, 66, 68, 68, 74, 81, 87, 91}

	Sort(list)

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}

	list = make([]int, 1000)
	expectedList = make([]int, 1000)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(300)
		expectedList[i] = list[i]
	}
	sort.Ints(expectedList)

	Sort(list)

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}

	list = make([]int, 20000)
	expectedList = make([]int, 20000)
	for i := 0; i < len(list); i++ {
		list[i] = rand.Intn(2000)
		expectedList[i] = list[i]
	}
	sort.Ints(expectedList)

	Sort(list)

	if !reflect.DeepEqual(list, expectedList) {
		t.Error("Expected", expectedList, ", got ", list)
	}
}
