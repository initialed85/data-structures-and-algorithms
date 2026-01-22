package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindLongestSubstringWithoutDuplicateCharacters(t *testing.T) {
	require.Equal(
		t,
		"wke",
		FindLongestSubstringWithoutDuplicateCharacters("pwwkew"),
	)
}

func TestLRUCache(t *testing.T) {
	lruCache := NewLRUCache(2)

	lruCache.Set(1, 1)

	lruCache.Set(2, 2)

	require.Equal(t, 1, lruCache.Get(1))

	lruCache.Set(3, 3)

	require.Equal(t, -1, lruCache.Get(2))

	lruCache.Set(4, 4)

	require.Equal(t, -1, lruCache.Get(1))

	require.Equal(t, 3, lruCache.Get(3))

	require.Equal(t, 4, lruCache.Get(4))

	lruCache.Set(5, 5)
	lruCache.Set(5, 6)
	lruCache.Set(5, 7)
	lruCache.Set(5, 8)

	require.Equal(t, -1, lruCache.Get(4))
	require.Equal(t, -1, lruCache.Get(4))
	require.Equal(t, -1, lruCache.Get(4))
	require.Equal(t, 8, lruCache.Get(5))
}

func TestShortestPathAlgorithmForSocialNetwork(t *testing.T) {
	person1 := "Kali"
	person2 := "Joey"
	person3 := "Molly"
	person4 := "Rudy"
	person5 := "Nancy"
	person6 := "Alan"
	person7 := "Kayla"

	friendships := map[string][]string{
		person1: {person2},
		person2: {person3, person4},
		person3: {person4},
		person4: {person5, person6},
		person5: {person6},
		person6: {person7},
		person7: {},
	}

	require.Equal(
		t,
		[]string{"Kayla", "Alan", "Rudy", "Joey", "Kali"},
		ShortestPathAlgorithmForSocialNetwork(
			friendships,
			person1,
			person7,
		),
	)
}
