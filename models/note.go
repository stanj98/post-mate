package types

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"uid,omitempty"`
	Title string          `bson:"name" json:"title"`
	ContentBody string    `bson:"body" json:"body"`
	ContentImgURL string  `bson:"imgURL" json:"imgURL,omitempty"`
	CreatedDate time.Time `bson:"created_date" json:"created_date,omitempty"`
	UpdatedDate time.Time `bson:"updated_date" json:"updated_date,omitempty"`
}