package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Tasks struct {
	Tasks []TaskObj `json:"tasks"`
}

type TaskObj struct {
	Id     int    `json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

func createTask(id int, task, status string) TaskObj {
	data := TaskObj{
		id,
		task,
		status,
	}
	return data
}

func AddTask(taskValue, taskStatus string) {
	// Create slice to hold tasks
	tasks := []TaskObj{}

	// Open tasks JSON file
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened tasks.json")
	// Defer closing the tasks.json file
	defer jsonFile.Close()

	// Read JSON file as byte slice
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Convert byte slice to object
	json.Unmarshal(byteValue, &tasks)
	newId := len(tasks) + 1
	fmt.Println("Adding task", newId)

	// Create New Task
	singleTask := createTask(newId, taskValue, taskStatus)

	// Append new task to existing tasks
	tasks = append(tasks, singleTask)

	// Format tasks as JSON
	file, _ := json.MarshalIndent(tasks, "", "  ")

	// Write new tasks JSON file
	_ = ioutil.WriteFile("tasks.json", file, 0644)

	// for i := 0; i < len(tasks); i++ {
	// 	fmt.Println("=== Task", i, "===")
	// 	fmt.Println("ID: " + strconv.Itoa(tasks[i].Id))
	// 	fmt.Println("Task: " + tasks[i].Task)
	// 	fmt.Println("Status: " + tasks[i].Status)
	// }
}

func ReadTasks() {
	// Create slice to hold tasks
	tasks := []TaskObj{}

	// Open tasks JSON file
	jsonFile, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened tasks.json")
	// Defer closing the tasks.json file
	defer jsonFile.Close()

	// Read JSON file as byte slice
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Convert byte slice to object
	json.Unmarshal(byteValue, &tasks)

	for i := 0; i < len(tasks); i++ {
		fmt.Println("=== Task", i, "===")
		fmt.Println("ID: " + strconv.Itoa(tasks[i].Id))
		fmt.Println("Task: " + tasks[i].Task)
		fmt.Println("Status: " + tasks[i].Status)
	}

}
