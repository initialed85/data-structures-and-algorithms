package pkg

/*
- Hash Set (using builtins)
  - Remove duplicates from array
*/

type EasyHashSet struct {
	values map[int]struct{}
}

func NewEasyHashSet() *EasyHashSet {
	return &EasyHashSet{
		values: make(map[int]struct{}),
	}
}

func (s *EasyHashSet) Add(value int) {
	s.values[value] = struct{}{}
}

func (s *EasyHashSet) Exists(value int) bool {
	_, ok := s.values[value]
	return ok
}

func (s *EasyHashSet) Items() []int {
	items := make([]int, 0)

	for value := range s.values {
		items = append(items, value)
	}

	return items
}

func EasyRemoveDuplicatesFromArray(array [1024]int) []int {
	hashSet := NewEasyHashSet()

	for _, x := range array {
		hashSet.Add(x)
	}

	return hashSet.Items()
}
