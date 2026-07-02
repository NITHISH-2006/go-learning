package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mergeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

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

	json.NewEncoder(w).Encode(slice)

}

func main() {
	http.HandleFunc("/merge", mergeHandler)

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)
}
