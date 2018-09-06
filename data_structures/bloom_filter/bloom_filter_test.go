package bloom_filter

import (
	"testing"
)

func assetEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected `%v`. Got `%v`\n", expected, actual)
	}
}

func TestNewBloomFilter(t *testing.T) {
	b := NewBloomFilter(0.1, 2)
	assetEqual(t, 10, b.GetSize())
	assetEqual(t, 2, b.GetCapacity())

	b = NewBloomFilter(0.1, 5)
	assetEqual(t, 24, b.GetSize())
	assetEqual(t, 5, b.GetCapacity())

	b = NewBloomFilter(0.1, 10)
	assetEqual(t, 48, b.GetSize())
	assetEqual(t, 10, b.GetCapacity())

	b = NewBloomFilter(0.1, 100)
	assetEqual(t, 480, b.GetSize())
	assetEqual(t, 100, b.GetCapacity())

	b = NewBloomFilter(0.1, 1000)
	assetEqual(t, 4793, b.GetSize())
	assetEqual(t, 1000, b.GetCapacity())

	b = NewBloomFilter(0.01, 2)
	assetEqual(t, 20, b.GetSize())
	assetEqual(t, 2, b.GetCapacity())

	b = NewBloomFilter(0.01, 5)
	assetEqual(t, 48, b.GetSize())
	assetEqual(t, 5, b.GetCapacity())

	b = NewBloomFilter(0.01, 10)
	assetEqual(t, 96, b.GetSize())
	assetEqual(t, 10, b.GetCapacity())

	b = NewBloomFilter(0.01, 100)
	assetEqual(t, 959, b.GetSize())
	assetEqual(t, 100, b.GetCapacity())

	b = NewBloomFilter(0.01, 1000)
	assetEqual(t, 9586, b.GetSize())
	assetEqual(t, 1000, b.GetCapacity())

	b = NewBloomFilter(0.0005, 2)
	assetEqual(t, 32, b.GetSize())
	assetEqual(t, 2, b.GetCapacity())

	b = NewBloomFilter(0.0005, 5)
	assetEqual(t, 80, b.GetSize())
	assetEqual(t, 5, b.GetCapacity())

	b = NewBloomFilter(0.0005, 10)
	assetEqual(t, 159, b.GetSize())
	assetEqual(t, 10, b.GetCapacity())

	b = NewBloomFilter(0.0005, 100)
	assetEqual(t, 1583, b.GetSize())
	assetEqual(t, 100, b.GetCapacity())

	b = NewBloomFilter(0.0005, 1000)
	assetEqual(t, 15821, b.GetSize())
	assetEqual(t, 1000, b.GetCapacity())

	b = NewBloomFilter(0.0005, 5000)
	assetEqual(t, 79102, b.GetSize())
	assetEqual(t, 5000, b.GetCapacity())
}

func TestSetBloomFilter(t *testing.T) {
	b := NewBloomFilter(0.001, 10)

	assetEqual(t, false, b.Check("user:35"))
	assetEqual(t, false, b.Check("user:165"))
	assetEqual(t, false, b.Check("user:55"))
	assetEqual(t, false, b.Check("user:99"))

	b.Add("user:35")

	assetEqual(t, true, b.Check("user:35"))
	assetEqual(t, false, b.Check("user:165"))
	assetEqual(t, false, b.Check("user:55"))
	assetEqual(t, false, b.Check("user:99"))

	b.Add("user:165")

	assetEqual(t, true, b.Check("user:35"))
	assetEqual(t, true, b.Check("user:165"))
	assetEqual(t, false, b.Check("user:55"))
	assetEqual(t, false, b.Check("user:99"))

	b.Add("user:55")
	b.Add("user:99")

	assetEqual(t, true, b.Check("user:35"))
	assetEqual(t, true, b.Check("user:165"))
	assetEqual(t, true, b.Check("user:55"))
	assetEqual(t, true, b.Check("user:99"))
}
