package models

import (
	"time"
)

type Video struct {
	ID          int
	Title       string
	Description string
	Duration    string
	URL         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Annotation struct {
	ID        int
	VideoID   int
	StartTime string
	EndTime   string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
