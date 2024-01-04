package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stanj98/post-mate/types"
	"github.com/stanj98/post-mate/util"
)

var notes = []*types.Note {}

// var notes = []*types.Note {
// 	{
// 		Id: "1",
// 		Title: "First Note",
// 		ContentBody: "My first note",
// 		ContentImgURL: "https://png.pngtree.com/element_pic/12/03/20/1656e3fa305853d.jpg", 
// 		CreatedDate: time.Now(),
// 	},
// 	{
// 		Id: "3",
// 		Title: "Third Note",
// 		ContentBody: "My boring note", 
// 		CreatedDate: time.Now(),
// 	},
// } 

func GetNotesAPI(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data" : notes})
} 

func GetNoteAPI(c *gin.Context) {
	id := c.Param("id")
	note, err := util.GetNoteById(notes, id)
	
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Unable to find note!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data" : note})
}

func CreateNoteAPI(c *gin.Context) {
	var newNote *types.Note
	if err := c.BindJSON(&newNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : err})
		return
	}
	util.SetData(newNote)
	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusCreated, gin.H{"data" : newNote})
}

func EditNoteAPI(c *gin.Context) {
	var oldNote *types.Note
	id := c.Param("id")
	if err := c.BindJSON(&oldNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error parsing JSON data"})
		return
	}
	note, err := util.EditNoteById(notes, id, oldNote)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error! Unable to edit note"})
		return
	}
	c.IndentedJSON(http.StatusOK, note)
}

func DeleteNoteAPI(c *gin.Context) {
	id := c.Param("id")
	note, err := util.GetNoteById(notes, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error! Unable to find note"})
		return
	}
	notes, err := util.DeleteNote(notes, note)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Error! Unable to delete note"})
		return
	}
	data := gin.H{
		"notes": notes,
	}
	c.IndentedJSON(http.StatusOK, data)
}