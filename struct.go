package main

import (
	"fmt"
	"strings"
)

// User represents a person.
type User struct {
	Name string
	Age  int
}

// Greet returns a greeting message.
func (u User) Greet() string {
	return "Hi, " + u.Name + "!"
}

// Book represents a book.
type Book struct {
	Title  string
	Author string
	Pages  int
}

// Summary returns a formatted description of the book.
func (b Book) Summary() string {
	return fmt.Sprintf("%s by %s (%d pages)", b.Title, b.Author, b.Pages)
}

// wordCount counts the occurrences of each word in a sentence.
func wordCount(s string) map[string]int {
	words := strings.Fields(s)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

// removeDuplicates removes duplicate integers while preserving order.
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}

	return result
}

func main() {

	user := User{
		Name: "Nithish",
		Age:  20,
	}

	book := Book{
		Title:  "War and Peace",
		Author: "Tolstoy",
		Pages:  1225,
	}

	nums := []int{1, 2, 2, 3, 4, 5, 3, 6, 4}

	fmt.Println("===== User =====")
	fmt.Println(user.Greet())

	fmt.Println("\n===== Book =====")
	fmt.Println(book.Summary())

	fmt.Println("\n===== Word Count =====")
	fmt.Println(wordCount("go is fun go go fun"))

	fmt.Println("\n===== Remove Duplicates =====")
	fmt.Println(removeDuplicates(nums))
}