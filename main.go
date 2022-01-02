package main

import (
	"github.com/716green/go-todo/math"
	task "github.com/716green/go-todo/tasks"
	"github.com/716green/go-todo/utils"
)

func main() {
	sum := math.Add(1, 2)
	utils.PrintIntData("SUM", sum)
	// task.AddTask("Buy Groceries", "Pending")
	task.ReadTasks()

}
