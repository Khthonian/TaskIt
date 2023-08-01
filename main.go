package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
	// Iterate through the tasks array
	for index, task := range tasks {
		// Check if the task matches the ID or the hash value
		if (id != 0 && task.ID == id) || (hash != "" && task.Hash == hash) {
			taskName := tasks[index].Name

			// Slice before and after the index to exclude the deleted task
			tasks = append(tasks[:index], tasks[index+1:]...)

			// Alert the user
			successAlert := fmt.Sprintf("The task, '%s', was successfully deleted", taskName)
			fmt.Println(successAlert)

			return
		}
	}
	// Alert user of an erroneous input
	fmt.Println("The task could not be found.")
}

// Define a function to list the tasks
func listTask() {
	// Iterate through the tasks array
	for _, task := range tasks {
		// Default the status text to incomplete
		taskStatus := "Incomplete"
		// Check if the task is complete
		if task.Complete {
			// Change the status text to complete
			taskStatus = "Complete"
		}

		// Print the status of the task
		fmt.Printf("Task: %s\n ID: %d\n Status: %s\n", task.Name, task.ID, taskStatus)
	}
}

// Define a function to save tasks to a JSON file
func saveTasks() error {
	// Create a new file named tasks.json
	file, fail := os.Create("tasks.json")
	// If file creation fails, return error
	if fail != nil {
		return fail
	}

	// Close the file at the end
	defer file.Close()

	// Create a new JSON encoder to write to the file
	encoder := json.NewEncoder(file)
	// Encode the tasks in JSON formatting and write to the file
	fail = encoder.Encode(tasks)
	// Return any encoding errors
	return fail
}

// Define a function to load tasks from a JSON file
func loadTasks() error {
	// Open the JSON file
	file, fail := os.Open("tasks.json")
	// If file creation fails, return error
	if fail != nil {
		return fail
	}

	// Close the file at the end
	defer file.Close()

	// Create a new JSON decoder to read from the file
	decoder := json.NewDecoder(file)
	// Decode the tasks from JSON formatting and write to the tasks slice
	fail = decoder.Decode(&tasks)
	// Return any decoding errors
	return fail
}

func main() {
	// Load the tasks
	failLoad := loadTasks()
	if failLoad != nil && os.IsNotExist(failLoad) {
		fmt.Println("Failed to load tasks:", failLoad)
	}

	// Define flags for CLI usage
	var process, taskName, taskHash string
	var taskID int

	flag.StringVar(&process, "p", "", "Enter the process.")
	flag.StringVar(&taskName, "t", "", "Enter the task name")
	flag.IntVar(&taskID, "id", 0, "Enter the task ID")
	flag.StringVar(&taskHash, "hash", "", "Enter the task hash value")

	// Parse command line arguments
	flag.Parse()

	switch process {
	case "create":
		createTask(taskName)
	case "complete":
		completeTask(taskID, taskHash)
	case "delete":
		deleteTask(taskID, taskHash)
	case "list":
		listTask()
	default:
		fmt.Println("Unknown process. Please specify a valid process using the -p flag.")
	}

	// Save the tasks
	failSave := saveTasks()
	if failSave != nil {
		fmt.Println("Failed to save tasks:", failSave)
	}
}
