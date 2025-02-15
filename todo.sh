#!/bin/bash

# Prompt the user to enter their name
NAME=$(gum input --placeholder "What's your name?")

# Find existing task files
EXISTING_FILES=($(ls *.txt 2>/dev/null))

if [[ ${#EXISTING_FILES[@]} -gt 0 ]]; then
    # Prompt the user to choose an existing file or create a new one
    CHOICE=$(gum choose --item.foreground 250 "Create New File" "${EXISTING_FILES[@]}")
    if [[ "$CHOICE" == "Create New File" ]]; then
        TASK_FILE=$(gum input --placeholder "Enter the name of your todo list")
        TASK_FILE="${TASK_FILE}.txt"
    else
        TASK_FILE="$CHOICE"
    fi
else
    # Prompt the user to create a new file
    TASK_FILE=$(gum input --placeholder "Enter the name of your todo list")
    TASK_FILE="${TASK_FILE}.txt"
fi

# Create the task file if it doesn't exist
touch "$TASK_FILE"

gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "Hello, $NAME! You are currently working on $(gum style --foreground 212 "$TASK_FILE")."

while true; do
    TASK=$(gum input --placeholder "What would you like to add to your todo list?")
    echo "$TASK" >> "$TASK_FILE"
    echo -e "You have added, $(gum style --foreground 212 "$TASK")."

    sleep 1; clear

    echo -e "Would you like to add another $(gum style --italic --foreground 99 'task')?\n"
    CHOICE=$(gum choose --item.foreground 250 "Yes" "No")

    [[ "$CHOICE" == "No" ]] && break
done

clear

echo -e "Would you like to read your $(gum style --italic --foreground 99 'todo list')?\n"
CHOICE=$(gum choose --item.foreground 250 "Yes" "No")

if [[ "$CHOICE" == "Yes" ]]; then
    echo -e "You have the following tasks:\n"
    TASKS=()
    while IFS= read -r line; do
        TASKS+=("$line")
    done < "$TASK_FILE"
    for i in "${!TASKS[@]}"; do
        echo "$((i+1)). $(gum style --foreground 212 "${TASKS[i]}")"
    done
else
    echo "Okay, I'll discard it."
fi

echo -e "Would you like to delete a task from your $(gum style --italic --foreground 99 'todo list')?\n"
CHOICE=$(gum choose --item.foreground 250 "Yes" "No")

if [[ "$CHOICE" == "Yes" ]]; then
    echo -e "Select the task number to delete:\n"
    DELETE_INDEX=$(gum input --placeholder "Enter task number to delete")
    DELETE_INDEX=$((DELETE_INDEX-1))
    if [[ $DELETE_INDEX -ge 0 && $DELETE_INDEX -lt ${#TASKS[@]} ]]; then
        unset TASKS[$DELETE_INDEX]
        echo "${TASKS[@]}" > "$TASK_FILE"
        echo "Task deleted."
    else
        echo "Invalid task number."
    fi
else
    echo "No task deleted."
fi

echo -e "Would you like to delete a task file?\n"
CHOICE=$(gum choose --item.foreground 250 "Yes" "No")

if [[ "$CHOICE" == "Yes" ]]; then
    FILE_TO_DELETE=$(gum input --placeholder "Enter the name of the file to delete (without .txt)")
    FILE_TO_DELETE="${FILE_TO_DELETE}.txt"
    if [[ -f "$FILE_TO_DELETE" ]]; then
        rm "$FILE_TO_DELETE"
        echo "File $FILE_TO_DELETE deleted."
    else
        echo "File $FILE_TO_DELETE does not exist."
    fi
else
    echo "No file deleted."
fi

READ="Read"; CLOSE="Close"
ACTIONS=$(gum choose --no-limit "$READ" "$CLOSE")

clear; echo "One moment, please."

grep -q "$READ" <<< "$ACTIONS" && gum spin -s line --title "Reading the todo list..." -- sleep 1
grep -q "$CLOSE" <<< "$ACTIONS" && gum spin -s monkey --title " Closing your todo list..." -- sleep 2

sleep 1; clear

gum spin --title "Thinking about$(gum style --foreground "#04B575" "$GUM") your tasks..." -- sleep 2
clear

NICE_MEETING_YOU=$(gum style --height 5 --width 20 --padding '1 3' --border double --border-foreground 57 "Thanks for updating me, $(gum style --foreground 212 "$NAME"). See you soon!")
gum join --horizontal "$NICE_MEETING_YOU"