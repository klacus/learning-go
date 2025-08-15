package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Title:")
	content := getUserInput("Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	// You can not use Scanln for multi work input!

	// Use the bufio package to read from input allowing the user to enter multi-word text.
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // you need to use '' here not ""!
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	// Remove the newline and return character from the input.
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r") // if you are on Windows...

	return text
}
