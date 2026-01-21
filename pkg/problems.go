package pkg

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
)

func FindLongestSubstringWithoutDuplicateCharacters(input string) string {
	var i = 0
	var j = 0

	longestSubstringWithoutDuplicates := ""

	var hashSet map[rune]struct{}

	for i < len(input) && j < len(input) {
		hashSet = make(map[rune]struct{})

		slice := input[i:j]
		if slice == "" {
			j += 1
			continue
		}

		for _, c := range slice {
			hashSet[c] = struct{}{}
		}

		if len(hashSet) != len(slice) {
			i += 1
			continue
		}

		if len(slice) > len(longestSubstringWithoutDuplicates) {
			longestSubstringWithoutDuplicates = slice
		}

		j += 1
	}

	return longestSubstringWithoutDuplicates
}

type LRUCache struct {
	capacity int
	length   int
	items    *DoublyLinkedList[int]
	lookup   map[int]int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		length:   0,
		items:    NewDoublyLinkedList[int](),
		lookup:   make(map[int]int),
	}
}

func (c *LRUCache) Set(key int, value int) {
	if c.items.Length() == c.capacity {
		leastRecentlyUsed := c.items.PopLeft()
		if leastRecentlyUsed == nil {
			panic(fmt.Sprintf("assertion failed: doubly-linked list item for %#+v unexpectedly nil", key))
		}

		delete(c.lookup, *leastRecentlyUsed)
	}

	c.items.Push(key)
	c.lookup[key] = value

	// log.Printf("%#+v | %#+v", c.items.Items(), c.lookup)
}

func (c *LRUCache) Get(key int) int {
	value, ok := c.lookup[key]
	if !ok {
		return -1
	}

	c.items.Push(key)

	if c.items.Length() >= c.capacity {
		leastRecentlyUsed := c.items.PopLeft()
		if leastRecentlyUsed == nil {
			panic(fmt.Sprintf("assertion failed: doubly-linked list item for %#+v unexpectedly nil", key))
		}

		if *leastRecentlyUsed != key {
			delete(c.lookup, *leastRecentlyUsed)
		}
	}

	// log.Printf("%#+v | %#+v", c.items.Items(), c.lookup)

	return value
}

type Person struct {
	Name              string
	DistanceFromStart int
	Friends           []*Person
}

func ShortestPathAlgorithmForSocialNetwork(friendships map[string][]string, sourceName string, destinationName string) []string {
	//
	// dump the friendship graph out so we can look at it
	//

	data := "digraph {\n"

	for personName, friendNames := range friendships {
		for _, friendName := range friendNames {
			data += fmt.Sprintf("%s -> %s\n", personName, friendName)
		}
	}

	data += "}\n"

	err := os.WriteFile("./social_graph.dot", []byte(data), 0o777)
	if err != nil {
		panic(err)
	}

	err = exec.Command("dot", "-Tpng", "social_graph.dot", "-o", "social_graph.png").Run()
	if err != nil {
		panic(err)
	}

	//
	// turn the input data into some objects we can use
	//

	personByName := make(map[string]*Person)

	for personName := range friendships {
		distanceFromStart := math.MaxInt
		if personName == sourceName {
			distanceFromStart = 0
		}

		personByName[personName] = &Person{
			Name:              personName,
			DistanceFromStart: distanceFromStart,
			Friends:           make([]*Person, 0),
		}
	}

	for personName, friendNames := range friendships {
		person := personByName[personName]

		for _, friendName := range friendNames {
			friend := personByName[friendName]

			found := false

			for _, existingFriend := range person.Friends {
				if existingFriend.Name == friendName {
					found = true
					break
				}
			}

			if !found {
				person.Friends = append(person.Friends, friend)
			}
		}
	}

	//
	// chuck everyone in the unvisited queue
	//

	unvisited := NewMinPriorityQueue[string, int]()

	for _, person := range personByName {
		unvisited.Push(person.Name, person.DistanceFromStart)
	}

	for len(*unvisited) > 0 {
		personName, distanceFromStart, err := unvisited.Pop()
		if err != nil {
			panic(err)
		}

		person, _ := personByName[personName]

		for _, friend := range person.Friends {
			if friend.DistanceFromStart == math.MaxInt {
				friend.DistanceFromStart = distanceFromStart + 1
			}

			log.Printf("friend %s of %s is %d from start", friend.Name, person.Name, friend.DistanceFromStart)
			unvisited.Adjust(friend.Name, friend.DistanceFromStart)
		}
	}

	//
	// dump the friendship graph with distances out so we can look at it
	//

	data = "digraph {\n"

	for _, person := range personByName {
		for _, friend := range person.Friends {
			data += fmt.Sprintf("%s_%d -> %s_%d\n", person.Name, person.DistanceFromStart, friend.Name, friend.DistanceFromStart)
		}
	}

	data += "}\n"

	err = os.WriteFile("./social_graph_with_distances.dot", []byte(data), 0o777)
	if err != nil {
		panic(err)
	}

	err = exec.Command("dot", "-Tpng", "social_graph_with_distances.dot", "-o", "social_graph_with_distances.png").Run()
	if err != nil {
		panic(err)
	}

	//
	// propagate the friends
	//

	for _, person := range personByName {
		for _, friend := range person.Friends {
			found := false

			for _, otherFriend := range friend.Friends {
				if otherFriend.Name == person.Name {
					found = true
					break
				}
			}

			if !found {
				friend.Friends = append(friend.Friends, person)
			}
		}
	}

	//
	// now iterate from destination back to source
	//

	destination := personByName[destinationName]

	var current = destination

	shortestPath := make([]string, 0)

	totalDistanceFromStart := 0

	for current.Name != sourceName {
		unvisited = NewMinPriorityQueue[string, int]()

		for _, friend := range current.Friends {
			unvisited.Push(friend.Name, friend.DistanceFromStart)
		}

		currentName, distanceFromStart, err := unvisited.Pop()
		if err != nil {
			break
		}

		totalDistanceFromStart += distanceFromStart

		current = personByName[currentName]

		shortestPath = append(shortestPath, current.Name)
	}

	log.Printf("%d via %#+v", totalDistanceFromStart, shortestPath)

	return shortestPath
}
