package pkg

import (
	"fmt"
	"hash/maphash"
)

/*
- Hash Set (from scratch)
  - Remove duplicates from array
*/

var hardHashMapBuckets = 1024

type HardItem struct {
	Key   string
	Value int
}

type HardEntry struct {
	hash  uint64
	key   string
	value int
}

type HardHashMap struct {
	hasher  *maphash.Hash
	buckets [][]*HardEntry
}

func NewHardHashMap() *HardHashMap {
	hasher := &maphash.Hash{}
	hasher.SetSeed(maphash.MakeSeed())

	buckets := make([][]*HardEntry, hardHashMapBuckets)
	for i := range len(buckets) {
		buckets[i] = make([]*HardEntry, 0)
	}

	return &HardHashMap{
		hasher:  hasher,
		buckets: buckets,
	}
}

func (s *HardHashMap) getHash(value string) uint64 {
	s.hasher.Reset()
	maphash.WriteComparable(s.hasher, value)
	return s.hasher.Sum64()
}

func (s *HardHashMap) getBucket(hash uint64) int {
	return int(hash % uint64(hardHashMapBuckets))
}

func (s *HardHashMap) get(bucket int, key string) (int, error) {
	for i := range len(s.buckets[bucket]) {
		if s.buckets[bucket][i].key == key {
			return s.buckets[bucket][i].value, nil
		}
	}

	return 0, fmt.Errorf("unknown key %s", key)
}

func (s *HardHashMap) Get(key string) (int, error) {
	bucket := s.getBucket(s.getHash(key))

	return s.get(bucket, key)
}

func (s *HardHashMap) Exists(key string) bool {
	_, err := s.Get(key)
	return err == nil
}

func (s *HardHashMap) Set(key string, value int) {
	hash := s.getHash(key)
	bucket := s.getBucket(hash)

	_, err := s.get(bucket, key)
	if err != nil {
		s.buckets[bucket] = append(
			s.buckets[bucket],
			&HardEntry{
				hash:  hash,
				key:   key,
				value: value,
			},
		)

		return
	}

	for i := range len(s.buckets[bucket]) {
		if s.buckets[bucket][i].key == key {
			s.buckets[bucket][i].value = value
		}
	}
}

func (s *HardHashMap) Items() []HardItem {
	items := make([]HardItem, 0)

	for i := range len(s.buckets) {
		for j := range len(s.buckets[i]) {
			items = append(items, HardItem{
				Key:   s.buckets[i][j].key,
				Value: s.buckets[i][j].value,
			})
		}
	}

	return items
}

func HardCountOccurrencesOfString(array [10]string) []HardItem {
	hashMap := NewHardHashMap()

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
