package hash_map

import (
	"testing"
	"reflect"
)

func assetEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected `%v`. Got `%v`\n", expected, actual)
	}
}

func TestNewHashMap(t *testing.T) {
	m := NewHashMap()
	m.set("en", "English")

	val, ok := m.get("en")

	assetEqual(t, true, ok)
	assetEqual(t, "English", val)
}

func TestHashMapSet(t *testing.T) {
	m := NewHashMap()
	assetEqual(t, 0, m.length)
	assetEqual(t, 2, m.capacity)

	m.set("en", "English")
	assetEqual(t, 1, m.length)
	assetEqual(t, 2, m.capacity)

	m.set("ua", "Ukrainian")
	assetEqual(t, 2, m.length)
	assetEqual(t, 2, m.capacity)

	m.set("ru", "Russian")
	assetEqual(t, 3, m.length)
	assetEqual(t, 4, m.capacity)

	m.set("sv", "Swedish")
	assetEqual(t, 4, m.length)
	assetEqual(t, 8, m.capacity)

	m.set("au", "Australian")
	assetEqual(t, 5, m.length)
	assetEqual(t, 8, m.capacity)

	m.set("us", "USA")
	assetEqual(t, 6, m.length)
	assetEqual(t, 16, m.capacity)

	m.set("pl", "Poland")
	assetEqual(t, 7, m.length)
	assetEqual(t, 16, m.capacity)

	m.set("it", "Italian")
	assetEqual(t, 8, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok := m.get("en")
	assetEqual(t, true, ok)
	assetEqual(t, "English", val)

	val, ok = m.get("ua")
	assetEqual(t, true, ok)
	assetEqual(t, "Ukrainian", val)

	val, ok = m.get("ru")
	assetEqual(t, true, ok)
	assetEqual(t, "Russian", val)

	val, ok = m.get("sv")
	assetEqual(t, true, ok)
	assetEqual(t, "Swedish", val)

	val, ok = m.get("it")
	assetEqual(t, true, ok)
	assetEqual(t, "Italian", val)

	val, ok = m.get("pl")
	assetEqual(t, true, ok)
	assetEqual(t, "Poland", val)
}

func TestHashMapGet(t *testing.T) {
	m := NewHashMap()
	m.set("en", "English")
	m.set("ua", "Ukrainian")
	m.set("ru", "Russian")
	m.set("sv", "Swedish")
	m.set("au", "Australian")
	m.set("us", "USA")
	m.set("pl", "Poland")
	m.set("it", "Italian")

	m.set("en", "English (USA)")

	val, ok := m.get("en")
	assetEqual(t, true, ok)
	assetEqual(t, "English (USA)", val)

	val, ok = m.get("ua")
	assetEqual(t, true, ok)
	assetEqual(t, "Ukrainian", val)

	val, ok = m.get("ru")
	assetEqual(t, true, ok)
	assetEqual(t, "Russian", val)

	val, ok = m.get("sv")
	assetEqual(t, true, ok)
	assetEqual(t, "Swedish", val)

	val, ok = m.get("it")
	assetEqual(t, true, ok)
	assetEqual(t, "Italian", val)

	val, ok = m.get("pl")
	assetEqual(t, true, ok)
	assetEqual(t, "Poland", val)

	val, ok = m.get("es")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
}

func TestHashMapDelete(t *testing.T) {
	m := NewHashMap()

	m.set("en", "English")
	m.set("ua", "Ukrainian")
	m.set("ru", "Russian")
	m.set("sv", "Swedish")
	m.set("au", "Australian")
	m.set("us", "USA")
	m.set("pl", "Poland")
	m.set("it", "Italian")

	val, ok := m.get("ru")
	assetEqual(t, true, ok)
	assetEqual(t, "Russian", val)

	ok = m.del("ru")
	assetEqual(t, true, ok)

	val, ok = m.get("ru")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 7, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok = m.get("en")
	assetEqual(t, true, ok)
	assetEqual(t, "English", val)

	ok = m.del("en")
	assetEqual(t, true, ok)

	val, ok = m.get("en")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 6, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok = m.get("au")
	assetEqual(t, true, ok)
	assetEqual(t, "Australian", val)

	ok = m.del("au")
	assetEqual(t, true, ok)

	val, ok = m.get("au")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 5, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok = m.get("it")
	assetEqual(t, true, ok)
	assetEqual(t, "Italian", val)

	ok = m.del("it")
	assetEqual(t, true, ok)

	val, ok = m.get("it")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 4, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok = m.get("pl")
	assetEqual(t, true, ok)
	assetEqual(t, "Poland", val)

	ok = m.del("pl")
	assetEqual(t, true, ok)

	val, ok = m.get("pl")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 3, m.length)
	assetEqual(t, 16, m.capacity)

	val, ok = m.get("ua")
	assetEqual(t, true, ok)
	assetEqual(t, "Ukrainian", val)

	ok = m.del("ua")
	assetEqual(t, true, ok)

	val, ok = m.get("ua")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 2, m.length)
	assetEqual(t, 8, m.capacity)

	val, ok = m.get("sv")
	assetEqual(t, true, ok)
	assetEqual(t, "Swedish", val)

	ok = m.del("sv")
	assetEqual(t, true, ok)

	val, ok = m.get("sv")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 1, m.length)
	assetEqual(t, 8, m.capacity)

	val, ok = m.get("us")
	assetEqual(t, true, ok)
	assetEqual(t, "USA", val)

	ok = m.del("us")
	assetEqual(t, true, ok)

	val, ok = m.get("us")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)
	assetEqual(t, 0, m.length)
	assetEqual(t, 4, m.capacity)
}

func TestHashMapDeleteNonexistentKey(t *testing.T) {
	m := NewHashMap()

	m.set("en", "English")
	m.set("ua", "Ukrainian")
	m.set("ru", "Russian")
	m.set("sv", "Swedish")
	m.set("au", "Australian")
	m.set("us", "USA")
	m.set("pl", "Poland")
	m.set("it", "Italian")

	val, ok := m.get("cz")
	assetEqual(t, false, ok)
	assetEqual(t, nil, val)

	ok = m.del("cz")
	assetEqual(t, false, ok)
}

func TestHashMapKeys(t *testing.T) {
	m := NewHashMap()

	m.set("en", "English")
	m.set("ua", "Ukrainian")

	keys := m.keys()
	assetEqual(t, 2, len(keys))

	for _, key := range []string {"en", "ua"} {
		assetEqual(t, true, contains(keys, key))
	}

	m.set("ru", "Russian")
	m.set("sv", "Swedish")
	m.set("au", "Australian")
	m.set("us", "USA")

	keys = m.keys()
	assetEqual(t, 6, len(keys))

	for _, key := range []string {"en", "ua", "ru", "sv", "au", "us"} {
		assetEqual(t, true, contains(keys, key))
	}

	m.set("pl", "Poland")
	m.set("it", "Italian")

	keys = m.keys()
	assetEqual(t, 8, len(keys))

	for _, key := range []string {"en", "ua", "ru", "sv", "au", "us", "pl", "it"} {
		assetEqual(t, true, contains(keys, key))
	}
}


func TestHashMapValues(t *testing.T) {
	m := NewHashMap()

	m.set("en", "English")
	m.set("ua", "Ukrainian")

	values := m.values()
	assetEqual(t, 2, len(values))

	for _, value := range []string {"English", "Ukrainian"} {
		assetEqual(t, true, contains(values, value))
	}

	m.set("ru", "Russian")
	m.set("sv", "Swedish")
	m.set("au", "Australian")
	m.set("us", "USA")

	values = m.values()
	assetEqual(t, 6, len(values))

	for _, value := range []string {"English", "Ukrainian", "Russian", "Swedish", "Australian", "USA"} {
		assetEqual(t, true, contains(values, value))
	}

	m.set("pl", "Poland")
	m.set("it", "Italian")

	values = m.values()
	assetEqual(t, 8, len(values))

	for _, value := range []string {"English", "Ukrainian", "Russian", "Swedish", "Australian", "USA", "Poland", "Italian"} {
		assetEqual(t, true, contains(values, value))
	}
}

func contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {
			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}
