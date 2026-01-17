package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// hint: find pkg | entr -n -r -cc -s "go test -v -bench=. ./pkg/"

func BenchmarkLinkedList(b *testing.B) {
	b.Run("FindMiddleOfSinglyLinkedList", func(b *testing.B) {
		inputHead := NewSinglyLinkedList(0).Add(1).Add(2).Add(3).Add(4).Add(5)

		middle := FindMiddleOfSinglyLinkedList(inputHead)

		require.Equal(
			b,
			3,
			middle,
		)

		for b.Loop() {
			_ = FindMiddleOfSinglyLinkedList(inputHead)
		}
	})

	b.Run("ReverseSinglyLinkedList", func(b *testing.B) {
		inputHead := NewSinglyLinkedList(0).Add(1).Add(2).Add(3).Add(4).Add(5)

		outputHead := ReverseSinglyLinkedList(inputHead)

		values := make([]int, 0)

		curr := outputHead

		for curr != nil {
			values = append(values, curr.value)

			curr = curr.next
		}

		require.Equal(
			b,
			[]int{5, 4, 3, 2, 1, 0},
			values,
		)

		for b.Loop() {
			_ = ReverseSinglyLinkedList(inputHead)
		}
	})

	b.Run("RotateSinglyLinkedList", func(b *testing.B) {
		inputHead := NewSinglyLinkedList(0).Add(1).Add(2).Add(3).Add(4).Add(5)

		outputHead := RotateSinglyLinkedList(inputHead)

		values := make([]int, 0)

		curr := outputHead

		for curr != nil {
			values = append(values, curr.value)

			curr = curr.next
		}

		require.Equal(
			b,
			[]int{5, 0, 1, 2, 3, 4},
			values,
		)

		inputHead = NewSinglyLinkedList(0).Add(1).Add(2).Add(3).Add(4).Add(5)

		for b.Loop() {
			inputHead = RotateSinglyLinkedList(inputHead)

			values := make([]int, 0)

			curr := outputHead

			for curr != nil {
				values = append(values, curr.value)

				curr = curr.next
			}

			require.Equal(
				b,
				[]int{5, 0, 1, 2, 3, 4},
				values,
			)
		}
	})
}
