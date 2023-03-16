package main

import (
	"fmt"
	"learn-go/ecommerce/api-gateway/src/database"
	"learn-go/ecommerce/api-gateway/src/pkgs/dotenv"
	"log"

	"github.com/google/uuid"
)

type Note struct {
	Id      uuid.UUID `json:"id,omitempty" gorm:"primary_key;column:id;type:uuid;default:gen_random_uuid();"`
	Title   string    `json:"title" gorm:"column:title;"`
	Content string    `json:"content" gorm:"column:content;"`
}

type NoteUpdate struct {
	Title *string `json:"title" gorm:"column:title;"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {
	dotenv.LoadTheEnv()

	db := database.CreateDbInstance()

	// Auto migrate
	db.AutoMigrate(&Note{})

	// Insert new note
	newNote := Note{Title: "Test 2", Content: "Test content 2"}
	if err := db.Create(&newNote); err != nil {
		fmt.Println(err)
	}

	// Find all by title
	var notes []Note
	db.Where("title = ?", "Test 2").Find(&notes)
	fmt.Println(notes)

	// Find one by id
	var note Note
	if err := db.Where("id = ?", "9ac4e966-a0f8-44d7-ad10-9cef0290109c").First(&note); err != nil {
		log.Println(err)
	}
	fmt.Println(note)

	// Update one by id
	newTitle := ""
	db.Table(Note{}.TableName()).Where("id = ?", "9ac4e966-a0f8-44d7-ad10-9cef0290109c").Updates(&NoteUpdate{Title: &newTitle})

	// Delete one by id
	// db.Table(Note{}.TableName()).Where("id = ?", "0323ab1f-b919-45a8-8c30-5659aca0b5a0").Delete(nil) // Hard delete
	// db.Table(Note{}.TableName()).Where("id = ?", "0323ab1f-b919-45a8-8c30-5659aca0b5a0").Update("deleted_at", "now()") // Soft delete

}
