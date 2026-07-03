package main

import "fmt"

// Small example demonstrating basic error handling patterns
func mayFail(ok bool) error {
	if !ok {
		return fmt.Errorf("something went wrong")
	}
	return nil
}

func main() {
	if err := mayFail(false); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("no error")
}
