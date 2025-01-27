package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func showHeader(currentList string) {
	fmt.Println(RunCommand(fmt.Sprintf(`gum style --border double --margin "1" --padding "1" --foreground 212 "🎀 TODO LIST 🎀\n📂 Current List: %s"`, currentList)))
}

func showDivider() {
	fmt.Println(RunCommand(`gum style --foreground 219 "────────────────────────────────" `))
}

func SwitchList(currentList string) {
	lists := GetSavedLists()

	if len(lists) == 0 {
		fmt.Println("🎀 No lists available!")
		return
	}

	choice := RunCommand(fmt.Sprintf(`gum choose --header "Choose a TODO list" %s`, strings.Join(lists, " ")))

	if choice != "" {
		currentList = choice
		fmt.Printf("📂 Switched to: %s\n", currentList)
	} else {
		fmt.Println("🚫 No valid selection.")
	}
}

func GetSavedLists() []string {
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

func RunCommand(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err.Error()
	}
	return strings.TrimSpace(string(output))
}
