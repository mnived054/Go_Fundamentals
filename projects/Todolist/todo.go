package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	TaskName  string
	completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{TaskName: task, completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Task Added")
}

func listTasks() {
	for i, task := range tasks {
		var status string = "n"
		if task.completed {
			status = "d"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.TaskName, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].completed = true
		fmt.Println("Task marked as complete")
	} else {
		fmt.Println("Invalid index")
	}
}

func editTask(index int, newString string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].TaskName = newString
		fmt.Println("Task edited successfully")
	} else {
		fmt.Println("Invalid index")
	}
}

func deleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println("Task deleted successfully")
	} else {
		fmt.Println("Invalid Index")
	}
}

func main() {
	var indexInput int
	var taskInput, newTaskinput string

	fmt.Println("Options")
	fmt.Println("1. Add Tasks")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mask as completed")
	fmt.Println("4. Edit Task")
	fmt.Println("5. Delete Task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter choice (1, 2, 3, 4, 5, 6):")
		scanner.Scan()
		input := scanner.Text()

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}
		switch choice {
		case 1:
			fmt.Print("Enter task:")
			scanner.Scan()
			taskInput = scanner.Text()
			addTask(taskInput)
		case 2:
			listTasks()
		case 3:
			fmt.Print("Enter index:")
			scanner.Scan()

			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)
		case 4:
			fmt.Print("Enter index:")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			fmt.Print("Enter task:")
			scanner.Scan()
			newTaskinput = scanner.Text()
			editTask(indexInput, newTaskinput)
		case 5:
			fmt.Print("Enter index:")
			scanner.Scan()
			indexInput, _ = strconv.Atoi(scanner.Text())
			deleteTask(indexInput)
		case 6:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")

		}
	}

}
