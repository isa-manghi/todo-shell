# 🎀 TODO List 🎀

A visually appealing command-line TODO list manager built with **Golang** and **Gum**!

Manage multiple lists, track your tasks, and stay productive! 💖

## ✨ Features

✅ **Create & Manage Multiple Lists** (e.g., Work, Shopping, Personal)  
📂 **Auto-detect Saved Lists** for easy access  
🌸 **Cutesy UI with Colors, Borders, and Emojis**  
🎉 **Mark Tasks as Done** and track completed items  
🗑️ **Delete Tasks Easily**  
💖 **Save & Load Tasks Automatically**  
🎀 **Pretty Animations with `gum spin`**

---

## 📦 Installation

### 1️⃣ **Install Go**
Make sure you have Go installed. If not, download it from:  
👉 [https://go.dev/dl/](https://go.dev/dl/)

### 2️⃣ **Install Gum**
Gum is a tool for creating **cute terminal interfaces**! Install it using:
```sh
brew install gum   # macOS (Homebrew)
sudo pacman -S gum # Arch Linux
```

## 3️⃣ Clone this repository
```sh
git clone https://github.com/isa-manghi/todo.git
cd todo
```

## 4️⃣ Run the TODO Tool
```sh
go run main.go
```

## 🎀 Usage
📂 Choose or Create a List
When you start the tool, you'll be asked to choose a TODO list or create a new one (e.g., Work, Groceries, Personal).

## 📝 Main Menu
You'll see this menu:

``` sh
💖 Add Task
📜 List Tasks
✅ Mark Done
🗑️ Delete Task
🎀 View Completed
📂 Switch List
🚪 Exit
```

Simply pick an option and follow the prompts!

## 🎀 Example Usage
### 1️⃣ Add a Task
```console
💖 Enter a new task: Buy strawberries 🍓
🎀 Task added: Buy strawberries 🍓
```

### 2️⃣ List Tasks
```console
📜 TODO List:
🌸 1. Buy strawberries 🍓
🌸 2. Finish Golang project 💻
```

### 3️⃣ Mark Task as Done
```console 
✅ Pick a task to mark done: Buy strawberries 🍓
✨ Task completed: Buy strawberries 🍓
```

### 4️⃣ Switch Lists
```shell-session
📂 Pick a TODO list:
🆕 Create New List
📂 Work
📂 Groceries
📂 Personal
🎀 File Storage
```

Each list is stored in simple text files:

- todo_<list>.txt → Stores active tasks
- done_<list>.txt → Stores completed tasks

This makes it easy to back up or edit manually!

## 🎀 Future Ideas
- 💡 Priority Levels (⭐️ High, Medium, Low)
- 💡 Due Dates & Reminders

# 📜 License
This project is open-source under the MIT License. Feel free to contribute and improve it! 💖