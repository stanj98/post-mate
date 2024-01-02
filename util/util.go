package util

import (
	"github.com/google/uuid"
	"github.com/stanj98/post-mate/types"
	"errors"
	"time"
)

func generateUniqueId() string {
	return uuid.New().String()
}

func GetNoteById(notes []*types.Note, id string) (*types.Note, error) {
	for i, note := range notes {
		if note.Id == id {
			return notes[i], nil
		}
	}
	return nil, errors.New("Unable to find note!")
}

func EditNoteById(notes []*types.Note, id string, updatedNote *types.Note) (*types.Note, error) {
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

func SetData(note *types.Note) {
	note.Id = generateUniqueId()
	note.CreatedDate = time.Now()
}

func DeleteNote(notes []*types.Note, note *types.Note) ([]*types.Note, error) {
	for i, n := range notes {
		if n.Id == note.Id {
			notes = append(notes[:i], notes[i+1:]...)
			return notes, nil
		}
	}
	return nil, errors.New("Error! Cannot delete from list")
}