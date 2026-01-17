package pkg

import "fmt"

/*
- Hash Map (using builtins vs from scratch)
  - Count occurrence of strings in array
*/

type EasyItem struct {
	Key   string
	Value int
}

type EasyHashMap struct {
	items map[string]int
}

func NewEasyHashMap() *EasyHashMap {
	return &EasyHashMap{
		items: make(map[string]int),
	}
}

func (s *EasyHashMap) Set(key string, value int) {
	s.items[key] = value
}

func (s *EasyHashMap) Get(key string) (int, error) {
	value, ok := s.items[key]
	if !ok {
		return 0, fmt.Errorf("unknown key %s", key)
	}

	return value, nil
}

func (s *EasyHashMap) Exists(key string) bool {
	_, ok := s.items[key]
	return ok
}

func (s *EasyHashMap) Items() []EasyItem {
	items := make([]EasyItem, 0)

	for key, value := range s.items {
		items = append(items, EasyItem{
			Key:   key,
			Value: value,
		})
	}

	return items
}

func EasyCountOccurrencesOfString(array [10]string) []EasyItem {
	hashMap := NewEasyHashMap()

	for _, x := range array {
		count, err := hashMap.Get(x)
		if err != nil {
			hashMap.Set(x, 0)
		}

		count += 1

		hashMap.Set(x, count)
	}

	return hashMap.Items()
}
