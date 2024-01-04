package util

import (
	"github.com/google/uuid"
	"github.com/stanj98/post-mate/models"
	"errors"
	"time"
)

func generateUniqueId() string {
	return uuid.New().String()
}

func GetNoteById(notes []*models.Note, id string) (*models.Note, error) {
	for i, note := range notes {
		if note.Id == id {
			return notes[i], nil
		}
	}
	return nil, errors.New("Unable to find note!")
}

func EditNoteById(notes []*models.Note, id string, updatedNote *models.Note) (*models.Note, error) {
	note, err := GetNoteById(notes, id)
	if err != nil {
		return nil, err
	}
	note.Title = updatedNote.Title
	note.ContentBody = updatedNote.ContentBody
	note.ContentImgURL = updatedNote.ContentImgURL
	note.UpdatedDate = time.Now()
	return note, nil
}

func SetData(note *models.Note) {
	note.Id = generateUniqueId()
	note.CreatedDate = time.Now()
}

func DeleteNote(notes []*models.Note, note *models.Note) ([]*models.Note, error) {
	for i, n := range notes {
		if n.Id == note.Id {
			notes = append(notes[:i], notes[i+1:]...)
			return notes, nil
		}
	}
	return nil, errors.New("Error! Cannot delete from list")
}