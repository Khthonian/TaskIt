package main

import (
	"crypto/sha256"
	"fmt"
)

// Define the structure of a task
type Task struct {
	// Name of the task
	Name string
	// ID number of the task
	ID int
	// Additional hash value for the task
	Hash string
	// Has the task been completed?
	Complete bool
}

// Define an array to hold the tasks structures
var tasks []Task

// Define a function to create a hash
func generateHash(name string, id int) string {
	// Concatenate the name and ID
	input := fmt.Sprintf("%s%d", name, id)
	// Create a new SHA-256 hash object
	hasher := sha256.New()
	// Write the input to the hash object
	hasher.Write([]byte(input))
	// Return the hash
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// Define a function to create a task
func createTask(name string) {
	// Set the ID equal to one more than the current number of tasks
	id := len(tasks) + 1
	// Generate a hash from the name and ID
	hash := generateHash(name, id)

	// Define the new task
	task := Task{
		Name:     name,
		ID:       id,
		Hash:     hash,
		Complete: false,
	}

	// Add the new task to the tasks array
	tasks = append(tasks, task)

	// Alert the user
	successAlert := fmt.Sprintf("The task, '%s', was successfully created", name)
	fmt.Println(successAlert)
}

// Define a function to complete a task
func completeTask(id int, hash string) {
	// Iterate through the tasks array
	for index, task := range tasks {
		// Check if the task matches the ID or the hash value
		if (id != 0 && task.ID == id) || (hash != "" && task.Hash == hash) {
			// Switch the complete bool value to true
			tasks[index].Complete = true
			// Alert the user
			taskName := tasks[index].Name
			successAlert := fmt.Sprintf("The task, '%s', was successfully completed", taskName)
			fmt.Println(successAlert)
			return
		}
	}
	// Alert user of an erroneous input
	fmt.Println("The task could not be found.")
}

// Define a function to delete a task
func deleteTask(id int, hash string) {
	// Iterate though the tasks array
	for index, task := range tasks {
		// Check if the task matches the ID or the hash value
		if (id != 0 && task.ID == id) || (hash != "" && task.Hash == hash) {
			// Slice before and after the index to exclude the deleted task
			tasks = append(tasks[:index], tasks[index+1:]...)
			// Alert the user
			taskName := tasks[index].Name
			successAlert := fmt.Sprintf("The task, '%s', was successfully deleted", taskName)
			fmt.Println(successAlert)
			return
		}
	}
	// Alert user of an erroneous input
	fmt.Println("The task could not be found.")
}
