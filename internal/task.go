package internal

import (
	"fmt"
	"os"
	"strings"
)

func addTask(list string) {
	task := RunCommand(`gum input --placeholder "💖 Enter a new task" --prompt "✨ "`)
	if task == "" {
		fmt.Println("🚫 No task entered.")
		return
	}

	fmt.Println(RunCommand(`gum spin --spinner dot --title "Saving task..." -- sleep 1`))
	appendToFile(getTodoFile(list), task)
	fmt.Println("🎀 Task added: ", task)
}

func deleteTask(list string) {
	tasks := ListTasks(list)
	if len(tasks) == 0 {
		return
	}

	selectedTask := RunCommand(fmt.Sprintf(`gum choose --header "🗑️ Pick a task to delete!" %s`, strings.Join(tasks, " ")))
	if selectedTask == "" {
		fmt.Println("🚫 No task selected.")
		return
	}

	fmt.Println(RunCommand(`gum spin --spinner line --title "Deleting task..." -- sleep 1`))
	updateFile(getTodoFile(list), selectedTask, false)
	fmt.Println("🗑️ Task deleted: ", selectedTask)
}

func markDone(list string) {
	tasks := ListTasks(list)
	if len(tasks) == 0 {
		return
	}

	selectedTask := RunCommand(fmt.Sprintf(`gum choose --header "🎉 Pick a task to mark done!" %s`, strings.Join(tasks, " ")))
	if selectedTask == "" {
		fmt.Println("🚫 No task selected.")
		return
	}

	fmt.Println(RunCommand(`gum spin --spinner monkey --title "Marking as done..." -- sleep 1`))
	updateFile(getTodoFile(list), selectedTask, false)
	appendToFile(getDoneFile(list), selectedTask)
	fmt.Println("✨ Task completed: ", selectedTask)
}

func ListTasks(list string) []string {
	return readTasks(getTodoFile(list))
}

func ListDoneTasks(list string) []string {
	return readTasks(getDoneFile(list))
}

func readTasks(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil || len(data) == 0 {
		fmt.Println("🌸 No tasks found in", file)
		return nil
	}

	tasks := strings.Split(strings.TrimSpace(string(data)), "\n")
	return tasks
}

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

func SaveTasks(list string, tasks []string) {
	file := getTodoFile(list)
	data := strings.Join(tasks, "\n")
	err := os.WriteFile(file, []byte(data), 0644)
	if err != nil {
		fmt.Println("❌ Error saving tasks:", err)
	}
}

func getTodoFile(currentList string) string {
	return "todo_" + currentList + ".txt"
}

func getDoneFile(currentList string) string {
	return "done_" + currentList + ".txt"
}
