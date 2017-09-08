package set


import (
	"testing"
)

func TestSet(t *testing.T) {
	mySet := NewSet()

	if mySet.Len() != 0 {
		t.Error("Expected", 0, ", got ", mySet.Len())
	}

	mySet = NewSet(0, 5, 10, 15, 20, 25, 777, 32)

	if mySet.Len() != 8 {
		t.Error("Expected", 8, ", got ", mySet.Len())
	}

	mySet = NewSet(0, 17, -45)

	if mySet.Len() != 3 {
		t.Error("Expected", 3, ", got ", mySet.Len())
	}
}

func TestSet_Add(t *testing.T) {
	mySet := NewSet()

	if mySet.Len() != 0 {
		t.Error("Expected", 0, ", got ", mySet.Len())
	}

	if mySet.IsPresent(5) != false {
		t.Error("Expected", false, ", got ", mySet.IsPresent(5))
	}

	mySet.Add(5)

	if mySet.IsPresent(5) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(5))
	}

	if mySet.Len() != 1 {
		t.Error("Expected", 1, ", got ", mySet.Len())
	}

	if mySet.IsPresent(9) != false {
		t.Error("Expected", false, ", got ", mySet.IsPresent(9))
	}

	mySet.Add(9)

	if mySet.IsPresent(9) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(9))
	}

	if mySet.Len() != 2 {
		t.Error("Expected", 2, ", got ", mySet.Len())
	}

	mySet.Add(15)

	if mySet.Len() != 3 {
		t.Error("Expected", 3, ", got ", mySet.Len())
	}
}

func TestSet_Delete(t *testing.T) {
	mySet := NewSet( 5, 10, 15, 777, 32)

	if mySet.Len() != 5 {
		t.Error("Expected", 5, ", got ", mySet.Len())
	}

	if mySet.IsPresent(5) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(5))
	}
	if mySet.IsPresent(10) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(10))
	}
	if mySet.IsPresent(15) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(15))
	}
	if mySet.IsPresent(777) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(777))
	}
	if mySet.IsPresent(32) != true {
		t.Error("Expected", true, ", got ", mySet.IsPresent(32))
	}
}


func TestSet_Intersection(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	resSet := set1.Intersection(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 2, ", got ", resSet.Len())
	}

	set1 = NewSet( 3, 10, 10, 11, 15, 16, 25)
	set2 = NewSet( 1, 2, 3, 4, 5, 6, 7, 8, 9)

	resSet = set1.Intersection(set2)

	if resSet.Len() != 1 {
		t.Error("Expected", 1, ", got ", resSet.Len())
	}
	if resSet.IsPresent(3) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(3))
	}

	set1 = NewSet( )
	set2 = NewSet( 1, 2, 3, 4, 5, 6, 7, 8, 9)

	resSet = set1.Intersection(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 0, ", got ", resSet.Len())
	}

	set1 = NewSet( 5, 8, 14, 1)
	set2 = NewSet()

	resSet = set1.Intersection(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 0, ", got ", resSet.Len())
	}

	set1 = NewSet( 10, 15, 777)
	set2 = NewSet( 10, 215, 777, 24)

	resSet = set1.Intersection(set2)

	if resSet.Len() != 2 {
		t.Error("Expected", 2, ", got ", resSet.Len())
	}
	if resSet.IsPresent(10) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(10))
	}
	if resSet.IsPresent(777) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(777))
	}
	if resSet.IsPresent(15) != false {
		t.Error("Expected", false, ", got ", resSet.IsPresent(15))
	}
	if resSet.IsPresent(215) != false {
		t.Error("Expected", false, ", got ", resSet.IsPresent(215))
	}
	if resSet.IsPresent(24) != false {
		t.Error("Expected", false, ", got ", resSet.IsPresent(24))
	}

}

func TestSet_Difference(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	resSet := set1.Difference(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 2, ", got ", resSet.Len())
	}

	set1 = NewSet( 3, 10, 5, 15)
	set2 = NewSet(  3, 4, 5, 6, 7, 8, 10)

	resSet = set1.Difference(set2)

	if resSet.Len() != 1 {
		t.Error("Expected", 1, ", got ", resSet.Len())
	}
	if resSet.IsPresent(15) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(15))
	}

	set1 = NewSet( )
	set2 = NewSet( 1, 2, 3, 4, 5, 6, 7, 8, 9)

	resSet = set1.Difference(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 0, ", got ", resSet.Len())
	}

	set1 = NewSet( 5, 8, 14, 1)
	set2 = NewSet()

	resSet = set1.Difference(set2)

	if resSet.Len() != 4 {
		t.Error("Expected", 4, ", got ", resSet.Len())
	}
	if resSet.IsPresent(5) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(5))
	}
	if resSet.IsPresent(8) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(8))
	}
	if resSet.IsPresent(14) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(14))
	}
	if resSet.IsPresent(1) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(1))
	}

	set1 = NewSet( 15, 777, 5)
	set2 = NewSet( 10, 215, 24, 5, 7)

	resSet = set1.Difference(set2)

	if resSet.Len() != 2 {
		t.Error("Expected", 2, ", got ", resSet.Len())
	}
	if resSet.IsPresent(15) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(15))
	}
	if resSet.IsPresent(777) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(777))
	}
}

func TestSet_Union(t *testing.T) {
	set1 := NewSet()
	set2 := NewSet()

	resSet := set1.Union(set2)

	if resSet.Len() != 0 {
		t.Error("Expected", 2, ", got ", resSet.Len())
	}

	set1 = NewSet( )
	set2 = NewSet(  3, 4, 5)

	resSet = set1.Union(set2)

	if resSet.Len() != 3 {
		t.Error("Expected", 3, ", got ", resSet.Len())
	}
	if resSet.IsPresent(3) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(3))
	}
	if resSet.IsPresent(4) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(4))
	}
	if resSet.IsPresent(5) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(5))
	}

	set1 = NewSet( 3, 8, 9)
	set2 = NewSet()

	resSet = set1.Union(set2)

	if resSet.Len() != 3 {
		t.Error("Expected", 3, ", got ", resSet.Len())
	}
	if resSet.IsPresent(3) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(3))
	}
	if resSet.IsPresent(8) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(8))
	}
	if resSet.IsPresent(9) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(9))
	}

	set1 = NewSet( 15, 777, 5)
	set2 = NewSet( 10, 5, 7)

	resSet = set1.Union(set2)

	if resSet.Len() != 5 {
		t.Error("Expected", 5, ", got ", resSet.Len())
	}
	if resSet.IsPresent(15) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(15))
	}
	if resSet.IsPresent(777) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(777))
	}
	if resSet.IsPresent(5) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(5))
	}
	if resSet.IsPresent(10) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(10))
	}
	if resSet.IsPresent(7) != true {
		t.Error("Expected", true, ", got ", resSet.IsPresent(7))
	}
}
