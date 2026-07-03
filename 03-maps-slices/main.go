package main

import (
	"fmt"
	"sort"
)

func main() {
	list1 := []int{5, 2, 8, 2, 10}
	list2 := []int{8, 1, 9, 5}
	var slice []int

	seen := make(map[int]bool)

	for _, num := range list2 {
		list1 = append(list1, num)
	}

	for _, num := range list1 {
		if !seen[num] {
			seen[num] = true
			slice = append(slice, num)
		}

	}

	sort.Ints(slice)
	fmt.Println(slice)

}
