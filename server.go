package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
	"errors"
)

type Note struct {
	Id string             `json:"uid"`
	Title string          `json:"title"`
	ContentBody string    `json:"body"`
	ContentImgURL string  `json:"imgURL,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	UpdatedDate time.Time `json:"updated_date,omitempty"`
}

var notes = []Note {
	{
		Id: "1",
		Title: "First Note",
		ContentBody: "My first note",
		ContentImgURL: "https://png.pngtree.com/element_pic/12/03/20/1656e3fa305853d.jpg", 
		CreatedDate: time.Now(),
	},
	{
		Id: "2",
		Title: "Second Note",
		ContentBody: "My second note",
		ContentImgURL: "https://png.pngtree.com/element_pic/12/03/20/1656e3fa305853d.jpg", 
		CreatedDate: time.Now(),
	},
	{
		Id: "3",
		Title: "Third Note",
		ContentBody: "My boring note", 
		CreatedDate: time.Now(),
	},
} 

func generateUniqueId() string {
	return uuid.New().String()
}

func SetData(note *Note) {
	note.Id = generateUniqueId()
	note.CreatedDate = time.Now()
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data" : notes})
} 

func GetNoteById(id string) (*Note, error) {
	for i, note := range notes {
		if note.Id == id {
			return &notes[i], nil
		}
	}
	return nil, errors.New("Unable to find note!")
}

func getNote(c *gin.Context) {
	id := c.Param("id")
	note, err := GetNoteById(id)
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Unable to find note!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data" : note})
}

func createNote(c *gin.Context) {
	var newNote Note
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err})
		return
	}
	SetData(&newNote)
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, gin.H{"data" : newNote})
}

func EditNoteById(id string, updatedNote *Note) (*Note, error) {
	note, err := GetNoteById(id)
	if err != nil {
		return nil, err
	}
	note.Title = updatedNote.Title
	note.ContentBody = updatedNote.ContentBody
	note.ContentImgURL = updatedNote.ContentImgURL
	note.UpdatedDate = time.Now()
	return note, nil
}

func editNote(c *gin.Context) {
	var oldNote Note
	id := c.Param("id")
	if err := c.BindJSON(&oldNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error parsing JSON data"})
		return
	}
	note, err := EditNoteById(id, &oldNote)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error! Unable to edit note"})
		return
	}
	c.IndentedJSON(http.StatusOK, note)
}

func DeleteNote(note *Note) ([]Note, error) {
	for i, n := range notes {
		if n.Id == note.Id {
			notes = append(notes[:i], notes[i+1:]...)
			return notes, nil
		}
	}
	return nil, errors.New("Error! Cannot delete from list")
}

func deleteNote(c *gin.Context) {
	id := c.Param("id")
	note, err := GetNoteById(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error! Unable to find note"})
		return
	}
	notes, err := DeleteNote(note)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error! Unable to delete note"})
		return
	}
	c.IndentedJSON(http.StatusOK, notes)
}

func main() {

	/*
		Plan:
		2. Set up templates to communicate with created APIs
		3. Use mongodb instead of memory data in this file
		4. Refer to this: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

	*/

	// app.Static("/", "./public") 

	router := gin.Default()
	router.GET("/notes", getNotes)
	router.GET("/notes/:id", getNote)
	router.POST("/notes", createNote)
	router.PUT("/notes/:id", editNote)
	router.DELETE("/notes/:id", deleteNote)
	router.Run("localhost:8080")
}