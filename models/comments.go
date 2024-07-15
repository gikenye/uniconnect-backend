package models

import (
	"time"
	"uniconnect/graph/model"

	uuid "github.com/satori/go.uuid"
)

type Comment struct {
	BusinessID uuid.UUID `gorm:"type:uuid"`
	UserID     uuid.UUID `gorm:"type:uuid"`
	Sender     string
	DateSent   time.Time
	Message    string
}

func (c Comment) CreateToGraphData() *model.Comment {
	return &model.Comment{
		BusinessID: c.BusinessID.String(),
		Sender:     c.Sender,
		Message:    c.Message,
		Date:       c.DateSent,
	}
}
