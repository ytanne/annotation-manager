package models

import (
	"time"
)

type Video struct {
	ID          int        `json:"-"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	URL         string     `json:"url"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}

type Annotation struct {
	ID        int        `json:"id"`
	VideoID   int        `json:"video_id"`
	StartTime int        `json:"start_time"`
	EndTime   int        `json:"end_time"`
	Type      string     `json:"type"`
	Text      string     `json:"text"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
