package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error retrieving CWD: %v", err)
	}
	return cwd
}

func todoMdExists(cwd string) bool {
	todoFP := filepath.Join(cwd, "todo.md")

	_, err := os.Stat(todoFP)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		log.Fatalf("Error checking for todo.md: %v", err)
	}
	return true

}
func createTodoFile(cwd string) {
	todoFP := filepath.Join(cwd, "todo.md")

	file, err := os.Create(todoFP)
	if err != nil {
		log.Fatalf("Error creating todo.md: %v", err)
	}
	defer file.Close()
	log.Printf("Created %s", todoFP)
}

// TODO: Add item
func addItem(cwd string, items string) {
	todoFP := filepath.Join(cwd, "todo.md")

	file, err := os.OpenFile(todoFP, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(items + "\n")
	if err != nil {
		fmt.Printf("Error adding item to file todo.md: %v", err)
	} else {
		fmt.Println("Item added successfully")
	}
}

// TODO: List items
func listItems(cwd string) {
	todoFP := filepath.Join(cwd, "todo.md")
	file, err := os.Open(todoFP)
	if err != nil {
		fmt.Printf("Error opening todo file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text(), '\n')
	}

}

// TODO: Todo Help
func todoHelp() {
	fmt.Println("Welcome to Todo Help")
	fmt.Println("Actions: ")
	fmt.Println("\t Add Item: todo add <item>")
	fmt.Println("\t List all items: todo list")
	fmt.Println("\t Check item status: todo status <unique item string>")
	fmt.Println("\t Mark item completed: todo done <unique item string>")
}

// TODO: Item Status

func checkStatus(cwd string, itemID string) {
	todoFP := filepath.Join(cwd, "todo.md")
	file, err := os.Open(todoFP)
	if err != nil {
		fmt.Printf("Error loading todo file, %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		todoItem := scanner.Text()
		if strings.HasPrefix(todoItem, itemID) {
			fmt.Println(todoItem)
		}
	}

}

// TODO: Mark Completed
func markCompleted(cwd string, itemID string) {
	todoFP := filepath.Join(cwd, "todo.md")
	file, err := os.OpenFile(todoFP, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error loading todo file, %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tempStr := ""

	for scanner.Scan() {
		todoItem := scanner.Text()
		if strings.HasPrefix(todoItem, itemID) {
			tempStr += "[DONE]: " + todoItem

		} else {
			tempStr += todoItem

		}
	}
	_, err = file.WriteString(tempStr)

	if err != nil {
		fmt.Printf("Error marking as done: %v", err)
	} else {
		fmt.Println("Item marked successfully")
	}

}
func main() {
	cwd := getCwd()
	if !todoMdExists(cwd) {
		createTodoFile(cwd)
	} else {
		log.Println("File todo.md found in current directory")
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) >= 3 {
			addItem(cwd, strings.Join(os.Args[2:], " "))
		} else {
			fmt.Println("usage: todo add <item>")
		}
	case "list":
		listItems(cwd)

	case "status":
		if len(os.Args) == 3 {
			checkStatus(cwd, os.Args[2])
		} else {
			fmt.Println("usage: todo status <unique substring identifying todo item>")
		}

	case "done":
		if len(os.Args) == 3 {
			markCompleted(cwd, os.Args[2])
		} else {
			fmt.Println("usage: todo done <unique substring identifying todo item>")
		}
	case "help":
		todoHelp()
	default:
		todoHelp()

	}

}
