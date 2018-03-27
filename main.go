package main

import (
	"fmt"
	"os"
)

func main() {
  // This line is indented with 2 spaces
	// This line and below is indented with a tab
	fmt.Printf("Hello World.")

	if false {
		if false {
			if false {
				if false {
					// This is some deeply nested code
					os.Exit(1)
				}
			}
		}
	}
}
