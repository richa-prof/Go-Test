// pkg/todofetcher.go
package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// Todo represents the structure of a TODO item
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const baseURL = "https://jsonplaceholder.typicode.com/todos/"

// FetchTodos fetches the specified number of even-numbered TODOs concurrently
func FetchTodos(numTodos int) ([]Todo, error) {
	var todos []Todo
	var wg sync.WaitGroup

	// Use a channel to collect the results
	resultChan := make(chan Todo, numTodos)

	// Fetch TODOs concurrently
	for i := 2; i <= numTodos*2; i += 2 {
		wg.Add(1)
		go fetchTodo(i, &wg, resultChan)
	}

	// Close the result channel once all goroutines are done
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for todo := range resultChan {
		todos = append(todos, todo)
	}

	return todos, nil
}

func fetchTodo(todoID int, wg *sync.WaitGroup, resultChan chan<- Todo) {
	defer wg.Done()

	url := fmt.Sprintf("%s%d", baseURL, todoID)
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching TODO %d: %v\n", todoID, err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body for TODO %d: %v\n", todoID, err)
		return
	}

	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON for TODO %d: %v\n", todoID, err)
		return
	}

	resultChan <- todo
}
