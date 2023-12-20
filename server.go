package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
	"errors"
	"encoding/json"
	"bytes"
	"fmt"
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

func ViewNotes(c *gin.Context) {
	data := gin.H{
		"notes" : notes,
	}
	c.HTML(http.StatusOK, "view-notes.html", data)
}

func ViewNote(c *gin.Context) {
	id := c.Param("id")
	note, err := GetNoteById(id)
	if err != nil {
		c.HTML(http.StatusOK, "create-note.html", nil)
	}
	
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "view-note.html", data)
}

func EditNote(c *gin.Context) {
	id := c.Param("id")
	note, err := GetNoteById(id)
	if err != nil {
		c.HTML(http.StatusOK, "view-note.html", nil)
	}
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "edit-note.html", data)
}

func CloneNote(c *gin.Context) {
	id := c.Param("id")
	note, err := GetNoteById(id)
	if err != nil {
		c.HTML(http.StatusOK, "view-note.html", nil)
	}
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "create-note.html", data)
}

func CreateNote(c *gin.Context) {
	//try getting data inputted from user when going from edit note -> create note button click and show in create note page
	//like how quicknote does

	title := c.PostForm("input-title")
	content := c.PostForm("content")
	if title != "" && content != "" {
		payload, err := json.Marshal(map[string]string{
			"title" : title,
			"body": content,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error" : "Failed to process request"})
			return
		}

		apiUrl := "http://localhost:8080/api/notes"

		resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error" : "Failed to call API"})
			return
		}

		defer resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(http.StatusCreated)
		
		if resp.StatusCode != http.StatusCreated {
			c.AbortWithStatusJSON(resp.StatusCode, gin.H{"error" : "Failed to create note on the server"})
			return
		}

		data := gin.H{
			"notes" : notes,
		}

		c.HTML(http.StatusOK, "view-notes.html", data)
		return
	}
	c.HTML(http.StatusOK, "create-note.html", nil)
}

func main() {

	/*
		Plan:
		2. Set up templates to communicate with created APIs
		3. Use mongodb instead of memory data in this file
		4. Refer to this: https://dev.to/percoguru/getting-started-with-apis-in-golang-feat-fiber-and-gorm-2n34

	*/

	router := gin.Default()
	router.Static("/static", "./templates")
	router.LoadHTMLGlob("templates/*.html")

	viewRoutes := router.Group("/") 
	{
		viewRoutes.GET("/view-notes", ViewNotes)
		viewRoutes.GET("/:id", ViewNote)
		viewRoutes.GET("/:id/edit", EditNote)
		viewRoutes.GET("/:id/clone", CloneNote)
		viewRoutes.GET("/create-note", CreateNote)
		viewRoutes.POST("/create-note", CreateNote)
	}

	apiRoutes := router.Group("/api") 
	{
		apiRoutes.GET("/notes", getNotes)
		apiRoutes.GET("/notes/:id", getNote)
		apiRoutes.POST("/notes", createNote)
		apiRoutes.PUT("/notes/:id", editNote)
		apiRoutes.DELETE("/notes/:id", deleteNote)
	}
	
	router.Run("localhost:8080")
}