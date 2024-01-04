package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/stanj98/post-mate/util"
)

func CreateNote(c *gin.Context) {
	//try getting data inputted from user when going from edit note -> create note button click and show in create note page
	//like how quicknote does
	c.HTML(http.StatusOK, "create-note.html", nil)
}

func CloneNote(c *gin.Context) {
	id := c.Param("id")
	note, err := util.GetNoteById(notes, id)
	if err != nil {
		c.HTML(http.StatusOK, "view-note.html", nil)
		return
	}
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "clone-note.html", data)
}


func EditNote(c *gin.Context) {
	id := c.Param("id")
	note, err := util.GetNoteById(notes, id)

	if err != nil {
		c.HTML(http.StatusOK, "view-note.html", nil)
		return
	}
	
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "edit-note.html", data)
}

func ViewNote(c *gin.Context) {
	id := c.Param("id")
	note, err := util.GetNoteById(notes, id)
	if err != nil {
		c.HTML(http.StatusOK, "create-note.html", nil)
		return
	}
	
	data := gin.H{
		"note": note,
	}
	c.HTML(http.StatusOK, "view-note.html", data)
}

func ViewNotes(c *gin.Context) {
	data := gin.H{
		"notes" : notes,
	}
	c.HTML(http.StatusOK, "view-notes.html", data)
}


