package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// JSON content usually have lower case field names, so we use struct tags like `json:"title"`.
// This adds metadata to the struct fields. The metadata key can be anything. How it is interpreted or used at all depends on the code that processes the metadata.
// Like in this case the json package looks for Json: tags and ignores others.
type Note struct {
	// we do not want these internal field available outside the struct
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title and content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")         // replace spaces with underscores for file name
	fileName = fmt.Sprintf("%v.json", strings.ToLower(fileName)) // convert to lower case for consistency
	// os.WriteFile(fmt.Sprintf("%v.txt", fileName), []byte(note.Content), 0644)

	// Use the json built in package to convert our struct to JSON format.
	// Important, that the json package does not export non public fields (i.e. with upper case)
	jsonData, err := json.Marshal(note)
	if err != nil {
		fmt.Println("Error marshalling note:", err)
		return err
	}

	return os.WriteFile(fmt.Sprintf("%v", fileName), jsonData, 0644)
}
