package bloom_filter

import (
	"testing"
)

func assetEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected `%v`. Got `%v`\n", expected, actual)
	}
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
