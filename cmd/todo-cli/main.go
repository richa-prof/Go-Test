// cmd/todo-cli/main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"todoapp/pkg"
)

func main() {
	// Parse command-line flags
	numTodos := flag.Int("n", 20, "Number of even-numbered TODOs to fetch")
	flag.Parse()

	// Ensure the number of TODOs to fetch is positive
	if *numTodos <= 0 {
		log.Fatal("Number of TODOs must be greater than 0")
	}

	// Fetch TODOs
	todos, err := pkg.FetchTodos(*numTodos)
	if err != nil {
		log.Fatalf("Error fetching TODOs: %v", err)
	}

	// Print results
	for _, todo := range todos {
		fmt.Printf("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
	}
}
