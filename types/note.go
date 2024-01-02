package types

import (
	"time"
)

type Note struct {
	Id string             `json:"uid"`
	Title string          `json:"title"`
	ContentBody string    `json:"body"`
	ContentImgURL string  `json:"imgURL,omitempty"`
	CreatedDate time.Time `json:"created_date,omitempty"`
	UpdatedDate time.Time `json:"updated_date,omitempty"`
}