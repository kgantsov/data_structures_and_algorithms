package set

type Set struct {
	items map[interface{}]bool
}

func NewSet(args ...interface{}) *Set {
	s := new(Set)
	s.items = make(map[interface{}]bool)

	for _, v := range args {
		s.Add(v)
	}

	return s
}

func (s *Set) Len() interface{} {
	return len(s.items)
}

func (s *Set) Add(value interface{}) {

	s.items[value] = true
}

func (s *Set) Delete(value interface{}) {
	delete(s.items, value)
}

func (s *Set) IsPresent(value interface{}) bool {
	_, ok := s.items[value]

	return ok
}

func (s *Set) Intersection(s1 *Set) *Set {
	newSet := new(Set)
	newSet.items = make(map[interface{}]bool)

	for k := range s.items {
		if s1.IsPresent(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

func (s *Set) Difference(s1 *Set) *Set {
	newSet := new(Set)
	newSet.items = make(map[interface{}]bool)

	for k := range s.items {
		if !s1.IsPresent(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

func (s *Set) Union(s1 *Set) *Set {
	newSet := new(Set)
	newSet.items = make(map[interface{}]bool)

	for k := range s.items {
		newSet.Add(k)
	}
	for k := range s1.items {
		newSet.Add(k)
	}

	return newSet
}

func (s *Set) values() []interface{} {
	var values []interface{}

	for k := range s.items {
		values = append(values, k)
	}

	return values
}
