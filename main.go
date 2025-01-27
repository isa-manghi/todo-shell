package main

import (
	"fmt"
	"github.com/charmbracelet/gum"
	"os"
	"os/exec"
	"strings"
)

var currentList string

// Runs a shell command and returns the output
func runCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(output))
}

// Gets the file names for the selected list
func getTodoFile() string {
	return fmt.Sprintf("todo_%s.txt", currentList)
}

func getDoneFile() string {
	return fmt.Sprintf("done_%s.txt", currentList)
}

// Shows a cute header
func showHeader() {
	fmt.Println(runCommand(fmt.Sprintf(`gum style --border double --margin "1" --padding "1" --foreground 212 "🎀 TODO LIST 🎀\n📂 Current List: %s"`, currentList)))
}

// Adds a new task
func addTask() {
	task := runCommand(`gum input --placeholder "💖 Enter a new task" --prompt "✨ "`)
	if task == "" {
		fmt.Println("🚫 No task entered.")
		return
	}

	fmt.Println(runCommand(`gum spin --spinner dot --title "Saving task..." -- sleep 1`))
	appendToFile(getTodoFile(), task)
	fmt.Println("🎀 Task added: ", task)
}

// Lists all pending tasks
func listTasks() []string {
	return readTasks(getTodoFile())
}

// Lists all completed tasks
func listDoneTasks() []string {
	return readTasks(getDoneFile())
}

// Reads tasks from a file
func readTasks(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil || len(data) == 0 {
		fmt.Println("🌸 No tasks found in", file)
		return nil
	}

	tasks := strings.Split(strings.TrimSpace(string(data)), "\n")
	return tasks
}

// Deletes a selected task
func deleteTask() {
	tasks := listTasks()
	if len(tasks) == 0 {
		return
	}

	selectedTask := runCommand(fmt.Sprintf(`gum choose --header "🗑️ Pick a task to delete!" %s`, strings.Join(tasks, " ")))
	if selectedTask == "" {
		fmt.Println("🚫 No task selected.")
		return
	}

	fmt.Println(runCommand(`gum spin --spinner line --title "Deleting task..." -- sleep 1`))
	updateFile(getTodoFile(), selectedTask, false)
	fmt.Println("🗑️ Task deleted: ", selectedTask)
}

// Marks a task as done
func markDone() {
	tasks := listTasks()
	if len(tasks) == 0 {
		return
	}

	selectedTask := runCommand(fmt.Sprintf(`gum choose --header "🎉 Pick a task to mark done!" %s`, strings.Join(tasks, " ")))
	if selectedTask == "" {
		fmt.Println("🚫 No task selected.")
		return
	}

	fmt.Println(runCommand(`gum spin --spinner monkey --title "Marking as done..." -- sleep 1`))
	updateFile(getTodoFile(), selectedTask, false)
	appendToFile(getDoneFile(), selectedTask)
	fmt.Println("✨ Task completed: ", selectedTask)
}

// Appends a task to a file
func appendToFile(file, task string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(task + "\n")
	if err != nil {
		fmt.Println("❌ Error writing to file:", err)
	}
}

// Updates a file by removing a task
func updateFile(file, task string, keep bool) {
	tasks := readTasks(file)
	var newTasks []string

	for _, t := range tasks {
		if t != task || keep {
			newTasks = append(newTasks, t)
		}
	}

	err := os.WriteFile(file, []byte(strings.Join(newTasks, "\n")), 0644)
	if err != nil {
		fmt.Println("❌ Error updating", file)
	}
}

// Switches between different lists
func switchList() {
	lists := getSavedLists()

	if len(lists) == 0 {
		fmt.Println("🎀 No lists available!")
		return
	}

	// Show the lists with a divider between options
	choice := gum.Choose("Choose a TODO list", lists...).ShowDivider(true).Prompt()

	if choice != "" {
		currentList = choice
		fmt.Printf("📂 Switched to: %s\n", currentList)
	} else {
		fmt.Println("🚫 No valid selection.")
	}
}

// Retrieves all saved lists by checking existing files
func getSavedLists() []string {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("🚫 Error reading directory:", err)
		return []string{}
	}

	var lists []string
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, "todo_") && strings.HasSuffix(name, ".txt") {
			listName := strings.TrimPrefix(name, "todo_")
			listName = strings.TrimSuffix(listName, ".txt")
			lists = append(lists, listName)
		}
	}

	if len(lists) == 0 {
		return []string{"default"} // If no lists are found, return a default list.
	}

	return lists
}

// Displays a section divider
func showDivider() {
	fmt.Println(runCommand(`gum style --foreground 219 "────────────────────────────────" `))
}

// Modify saveTasks() to create a new file if it doesn’t exist:
func saveTasks(tasks []string) {
	fileName := fmt.Sprintf("todo_%s.txt", currentList)

	// Ensure the file exists before writing
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Println("🛠 Creating new TODO file:", fileName)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("🚫 Error creating file:", err)
			return
		}
		file.Close() // Close immediately after creating
	}

	// Write tasks to file
	data := strings.Join(tasks, "\n")
	err := os.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		fmt.Println("🚫 Error saving tasks:", err)
	}
}

func main() {
	fmt.Println("🎀 TODO List is starting...")

	// Step 1: Check if any TODO files exist
	lists := getSavedLists()
	fmt.Println("🛠 Debug: Found lists →", lists)

	if len(lists) == 0 {
		// Step 2: If no lists are found, prompt for the first list
		fmt.Println("🚀 No TODO lists found! Let's create one now.")

		firstList := runCommand(`gum input --placeholder "Enter a name for your first list" --prompt "📂 "`)
		fmt.Println("🛠 Debug: User entered →", firstList)

		if firstList == "" {
			fmt.Println("🚫 No list name entered. Exiting.")
			return
		}

		// Step 3: Create the new list and save it
		currentList := strings.ReplaceAll(firstList, " ", "_") // No spaces allowed
		fmt.Println("🎀 Created list:", currentList)

		saveTasks([]string{}) // Create an empty task file for the new list
	} else {
		// If lists exist, let the user choose one
		fmt.Println("📂 Lists exist. Switching list...")
		switchList()
	}

	fmt.Println("🎀 Setup Complete! Program should continue now.")
}
