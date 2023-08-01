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
