package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	uiLoop()
}

func uiLoop() {
	isExit := 0
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Init")
	for isExit != 1 {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Add a new task")
		fmt.Println("2. Update task")
		fmt.Println("3. List tasks")
		fmt.Println("4. Delete task")
		fmt.Println("x. Exit")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		text = strings.TrimRight(text, "\n")
		switch text {
		case "1":
			fmt.Println("Enter new task:")
			newTask, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			createTask(newTask)
			fmt.Println("Task added")
		case "2":
			fmt.Println("Editing task")
		case "3":
			for _, line := range readTasks() {
				fmt.Println(line)
			}
		case "4":
			fmt.Println("Which task do you want to delete?")
			for i, line := range readTasks() {
				fmt.Printf("%d. %s\n", i, line)
			}
			if err != nil {
				panic(err)
			}

		case "x":
			fmt.Println("Exiting...")
			isExit = 1
		default:
			fmt.Println("IDK yet")
		}
	}
}

func createTask(taskName string) {
	currentTime := time.Now()
	task := Task{strings.TrimRight(taskName, "\n"), currentTime, false}

	f, err := os.OpenFile("tasks.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(task.String()); err != nil {
		panic(err)
	}
}

func readTasks() []string {
	content, err := os.ReadFile("tasks.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines
}

type Task struct {
	task      string
	createdAt time.Time
	isDone    bool
}

func (t Task) String() string {
	return fmt.Sprintln(t.task, t.createdAt, t.isDone)
}
