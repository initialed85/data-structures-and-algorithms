package pkg

/*
- Array
    - Linear search to find element
    - Fixed length sliding window to find sequence of known length
    - Variable length sliding window to find longest sequence of repeating pattern
*/

func LinearSearch(array [1024]int, value int) []int {
	/*
		Just basic array stuff, increment an index and check
		if the element at that index is the desired value
	*/

	indexes := make([]int, 0)

	for i := range 1024 {
		if array[i] == value {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func FixedLengthSlidingWindow(array [1024]int, sequence [8]int) (int, int) {
	/*
		Fairly straightforward, just maintain two indexes and
		advance them through the array at the same speed,
		slicing as you go
	*/

	start := -1
	stop := -1

	var slice [8]int
	_ = slice

	j := 8
	for i := range 1024 - 8 {
		slice = [8]int(array[i:j])
		if slice == sequence {
			start = i
			stop = j
			break
		}

		j += 1
	}

	return start, stop
}

func VariableLengthSlidingWindow(array [1024]int, pattern [2]int) (int, int) {
	/*
		Two indexes again, advance them forward together until
		the start of the pattern is seen, then keep advancing
		only the front index until the pattern is no longer seen

		Repeat each time the start of the pattern is seen and
		keep the length of the longest occurrence
	*/

	possibleStart := -1
	possibleStop := -1

	_ = possibleStart
	_ = possibleStop

	start := -1
	stop := -1

	inPattern := false

	i := 0
	for j := 1; j < 1024; j++ {
		if !inPattern {
			if array[i] == pattern[0] && array[j] == pattern[1] {
				inPattern = true
				possibleStart = i
				possibleStop = j
			}
		} else {
			if (array[j-1] == pattern[0] && array[j] == pattern[1]) ||
				(array[j-1] == pattern[1] && array[j] == pattern[0]) {
				possibleStop = j
			} else {
				inPattern = false
				if possibleStop-possibleStart > stop-start {
					start = possibleStart
					stop = possibleStop
				}
			}
		}

		i += 1
	}

	if start == -1 && stop == -1 {
		start = possibleStart
		stop = possibleStop
	}

	return start, stop
}
