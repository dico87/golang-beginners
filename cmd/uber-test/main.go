package main

import (
	"fmt"
	"sort"
)

func main() {
	chars := []string{"b", "a", "a", "b", "c"}
	top := 3
	occurs := getOccurs(chars)
	result := getTop(occurs, top)
	fmt.Printf("%v", result)
}

func getOccurs(chars []string) map[string]int {
	occurs := map[string]int{}
	for _, val := range chars {
		occurs[val] = occurs[val] + 1
	}
	return occurs
}

func getTop(occurs map[string]int, top int) []string {
	var keys []string
	for key := range occurs {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return occurs[keys[i]] > occurs[keys[j]]
	})

	return keys[:top]
}