package pkg

import (
	"hash/maphash"
)

var hardHashBuckets = 1024

type Entry struct {
	hash  uint64
	value int
}

type HardHashSet struct {
	hasher  *maphash.Hash
	buckets [][]*Entry
}

func NewHardHashSet() *HardHashSet {
	hasher := &maphash.Hash{}
	hasher.SetSeed(maphash.MakeSeed())

	buckets := make([][]*Entry, hardHashBuckets)
	for i := range len(buckets) {
		buckets[i] = make([]*Entry, 0)
	}

	return &HardHashSet{
		hasher:  hasher,
		buckets: buckets,
	}
}

func (s *HardHashSet) getHash(value int) uint64 {
	s.hasher.Reset()
	maphash.WriteComparable(s.hasher, value)
	return s.hasher.Sum64()
}

func (s *HardHashSet) getBucket(hash uint64) int {
	return int(hash % uint64(hardHashBuckets))
}

func (s *HardHashSet) exists(bucket int, value int) bool {
	for i := range len(s.buckets[bucket]) {
		if s.buckets[bucket][i].value == value {
			return true
		}
	}

	return false
}
func (s *HardHashSet) Exists(value int) bool {
	bucket := s.getBucket(s.getHash(value))

	return s.exists(bucket, value)
}

func (s *HardHashSet) Add(value int) {
	hash := s.getHash(value)
	bucket := s.getBucket(hash)

	if s.exists(bucket, value) {
		return
	}

	s.buckets[bucket] = append(
		s.buckets[bucket],
		&Entry{
			hash:  hash,
			value: value,
		},
	)
}

func (s *HardHashSet) Items() []int {
	items := make([]int, 0)

	for i := range len(s.buckets) {
		for j := range len(s.buckets[i]) {
			items = append(items, s.buckets[i][j].value)
		}
	}

	return items
}

func HardRemoveDuplicatesFromArray(array [1024]int) []int {
	hashSet := NewHardHashSet()

	for _, x := range array {
		hashSet.Add(x)
	}

	return hashSet.Items()
}
