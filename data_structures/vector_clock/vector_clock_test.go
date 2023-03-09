package vectorclock

import (
	"reflect"
	"testing"
)

func TestVectorClock(t *testing.T) {
	// Create a new VectorClock instance
	vc := NewVectorClock(3)

	// Test that GetClock() returns the correct initial clock value
	expected := []int64{0, 0, 0}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("GetClock() returned incorrect initial clock value: got %v, expected %v", vc.GetClock(), expected)
	}

	vc = NewVectorClock(5)

	// Test that GetClock() returns the correct initial clock value
	expected = []int64{0, 0, 0, 0, 0}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("GetClock() returned incorrect initial clock value: got %v, expected %v", vc.GetClock(), expected)
	}

}

func TestVectorClockIncrement(t *testing.T) {
	// Test that Increment() increments the correct node's clock value
	vc := NewVectorClock(4)

	vc.Increment(1)
	expected := []int64{0, 1, 0, 0}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Increment() did not increment the correct node's clock value: got %v, expected %v", vc.GetClock(), expected)
	}

	vc.Increment(1)
	expected = []int64{0, 2, 0, 0}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Increment() did not increment the correct node's clock value: got %v, expected %v", vc.GetClock(), expected)
	}

	vc.Increment(3)
	expected = []int64{0, 2, 0, 1}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Increment() did not increment the correct node's clock value: got %v, expected %v", vc.GetClock(), expected)
	}

	vc.Increment(0)
	expected = []int64{1, 2, 0, 1}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Increment() did not increment the correct node's clock value: got %v, expected %v", vc.GetClock(), expected)
	}
}

func TestVectorClockUpdate(t *testing.T) {
	// Test that Update() updates the clock correctly

	vc := NewVectorClock(5)

	vc.Update([]int64{0, 2, 4, 3, 7})
	expected := []int64{0, 2, 4, 3, 7}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}

	vc.Update([]int64{1, 0, 3, 5, 9})
	expected = []int64{1, 2, 4, 5, 9}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}

	vc.Update([]int64{1, 0, 3, 5, 9})
	expected = []int64{1, 2, 4, 5, 9}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}
}

func TestVectorClockCompare(t *testing.T) {
	vc := NewVectorClock(3)

	// Test that Update() updates the clock correctly
	vc.Update([]int64{0, 2, 0})
	expected := []int64{0, 2, 0}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}

	// Test that Compare() returns the correct comparison value
	// When the two clocks are equal
	result := vc.Compare([]int64{0, 2, 0})
	if result != 0 {
		t.Errorf("Compare() did not return the correct value when the clocks are equal: got %d, expected %d", result, 0)
	}

	// When the local clock is smaller than the other clock
	result = vc.Compare([]int64{0, 3, 0})
	if result != -1 {
		t.Errorf("Compare() did not return the correct value when local clock is smaller than the other clock: got %d, expected %d", result, -1)
	}

	// When the local clock is greater than the other clock
	result = vc.Compare([]int64{0, 1, 0})
	if result != 1 {
		t.Errorf("Compare() did not return the correct value when the local clock is greater than the other clock: got %d, expected %d", result, 1)
	}

	// Test that Update() updates the clock correctly
	vc.Update([]int64{1, 1, 3})
	expected = []int64{1, 2, 3}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}

	// Test that Compare() returns the correct comparison value
	// When the two clocks are equal
	result = vc.Compare([]int64{1, 2, 3})
	if result != 0 {
		t.Errorf("Compare() did not return the correct value when the clocks are equal: got %d, expected %d", result, 0)
	}

	// When the local clock is smaller than the other clock
	result = vc.Compare([]int64{2, 3, 4})
	if result != -1 {
		t.Errorf("Compare() did not return the correct value when local clock is smaller than the other clock: got %d, expected %d", result, -1)
	}

	// When the local clock is greater than the other clock
	result = vc.Compare([]int64{1, 2, 2})
	if result != 1 {
		t.Errorf("Compare() did not return the correct value when the local clock is greater than the other clock: got %d, expected %d", result, 1)
	}

	// Test that Update() updates the clock correctly
	vc.Update([]int64{3, 7, 2})
	expected = []int64{3, 7, 3}
	if !reflect.DeepEqual(vc.GetClock(), expected) {
		t.Errorf("Update() did not update the clock correctly: got %v, expected %v", vc.GetClock(), expected)
	}

	// Test that Compare() returns the correct comparison value
	// When the two clocks are equal
	result = vc.Compare([]int64{3, 7, 3})
	if result != 0 {
		t.Errorf("Compare() did not return the correct value when the clocks are equal: got %d, expected %d", result, 0)
	}

	// When the local clock is smaller than the other clock
	result = vc.Compare([]int64{3, 7, 4})
	if result != -1 {
		t.Errorf("Compare() did not return the correct value when local clock is smaller than the other clock: got %d, expected %d", result, -1)
	}

	result = vc.Compare([]int64{3, 9, 3})
	if result != -1 {
		t.Errorf("Compare() did not return the correct value when local clock is smaller than the other clock: got %d, expected %d", result, -1)
	}

	// When the local clock is greater than the other clock
	result = vc.Compare([]int64{3, 6, 3})
	if result != 1 {
		t.Errorf("Compare() did not return the correct value when the local clock is greater than the other clock: got %d, expected %d", result, 1)
	}

	result = vc.Compare([]int64{2, 7, 3})
	if result != 1 {
		t.Errorf("Compare() did not return the correct value when the local clock is greater than the other clock: got %d, expected %d", result, 1)
	}
}
