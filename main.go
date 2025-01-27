package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"todo/internal"
)

type model struct {
	tasks []string
	done  []string
}

func initialModel(currentList string) model {
	return model{
		tasks: internal.ListTasks(currentList),
		done:  internal.ListDoneTasks(currentList),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "TODO List:\n\n"
	for i, task := range m.tasks {
		s += fmt.Sprintf("%d. %s\n", i+1, task)
	}
	s += "\nPress q to quit.\n"
	return s
}

func main() {
	var currentList string
	fmt.Println("🎀 TODO List is starting...")

	lists := internal.GetSavedLists()
	fmt.Println("🛠 Debug: Found lists →", lists)

	if len(lists) == 0 {
		fmt.Println("🚀 No TODO lists found! Let's create one now.")

		firstList := internal.RunCommand(`gum input --placeholder "Enter a name for your first list" --prompt "📂 "`)
		fmt.Println("🛠 Debug: User entered →", firstList)

		if firstList == "" {
			fmt.Println("🚫 No list name entered. Exiting.")
			return
		}

		currentList = strings.ReplaceAll(firstList, " ", "_")
		fmt.Println("🎀 Created list:", currentList)

		internal.SaveTasks(currentList, []string{})
	} else {
		fmt.Println("📂 Lists exist. Switching list...")
		internal.SwitchList(currentList)
	}

	p := tea.NewProgram(initialModel(currentList))
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}
