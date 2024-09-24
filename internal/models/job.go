package models

import (
	"time"

	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"gorm.io/gorm"
)

const (
	PENDING    = "pending" // Starts from 1
	PROCESSING = "processing"
	COMPLETED  = "completed"
)

type Job struct {
	gorm.Model
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Link      *string   `json:"link" gorm:"index"`
	Status    string    `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (*Job) NewJob(link *string) Job {
	return Job{
		Id:        uuid.UUID{},
		Link:      link,
		Status:    PENDING,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
