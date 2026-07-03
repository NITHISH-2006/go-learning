package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

type cleanedData struct {
	CleanedIDs  []int  `json:"cleaned_ids"`
	TotalUnique int    `json:"total_unique"`
	Status      string `json:"status"`
}

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

func numbersHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	urls := r.URL.Query()["url"]

	if len(urls) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please provide urls"})
		return
	}

	json.NewEncoder(w).Encode(urls)
}

func cleanDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	slice := r.URL.Query()["id"]

	if len(slice) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Please provide id"})
		return
	}

	var cleaned []int

	seen := make(map[int]bool)

	for _, num := range slice {

		con, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		if !seen[con] {
			seen[con] = true
			cleaned = append(cleaned, con)
		}
	}

	sort.Ints(cleaned)

	response := cleanedData{
		CleanedIDs:  cleaned,
		TotalUnique: len(cleaned),
		Status:      "success",
	}

	json.NewEncoder(w).Encode(response)

}

func main() {

	http.HandleFunc("/merge", mergeHandler)
	http.HandleFunc("/number", numbersHandler)
	http.HandleFunc("/clean-data", cleanDataHandler)

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)

}
