package pkg

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

// hint: find pkg | entr -n -r -cc -s "go test -v -bench=. ./pkg/"

func BenchmarkHashMap(b *testing.B) {
	b.Run("EasyCountOccurrencesOfString", func(b *testing.B) {
		array := [10]string{"Apples", "Oranges", "Pears", "Pears", "Apples", "Apples", "Oranges", "Oranges", "Oranges", "Pears"}

		deduplicatedArray := EasyCountOccurrencesOfString(array)
		slices.SortFunc(deduplicatedArray, func(a EasyItem, b EasyItem) int {
			if a.Key < b.Key {
				return -1
			} else if a.Key > b.Key {
				return 1
			} else {
				return 0
			}
		})

		require.Equal(
			b,
			[]EasyItem{{Key: "Apples", Value: 3}, {Key: "Oranges", Value: 4}, {Key: "Pears", Value: 3}},
			deduplicatedArray,
		)

		for b.Loop() {
			_ = EasyCountOccurrencesOfString(array)
		}
	})

	b.Run("HardCountOccurrencesOfString", func(b *testing.B) {
		array := [10]string{"Apples", "Oranges", "Pears", "Pears", "Apples", "Apples", "Oranges", "Oranges", "Oranges", "Pears"}

		deduplicatedArray := HardCountOccurrencesOfString(array)
		slices.SortFunc(deduplicatedArray, func(a HardItem, b HardItem) int {
			if a.Key < b.Key {
				return -1
			} else if a.Key > b.Key {
				return 1
			} else {
				return 0
			}
		})

		require.Equal(
			b,
			[]HardItem{{Key: "Apples", Value: 3}, {Key: "Oranges", Value: 4}, {Key: "Pears", Value: 3}},
			deduplicatedArray,
		)

		for b.Loop() {
			_ = HardCountOccurrencesOfString(array)
		}
	})
}
