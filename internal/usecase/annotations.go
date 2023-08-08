package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ytanne/annotation-manager/pkg/models"
)

func (u usecase) AddAnnotation(ctx context.Context, annotation models.Annotation) (int, error) {
	st, err := time.ParseDuration(annotation.StartTime)
	if err != nil {
		return -1, fmt.Errorf("start time is invalid - %s", st)
	}

	et, err := time.ParseDuration(annotation.EndTime)
	if err != nil {
		return -1, fmt.Errorf("end time is invalid - %s", annotation.EndTime)
	}

	if st >= et {
		return -1, fmt.Errorf("start time should be greater than end time")
	}

	video, err := u.GetVideo(ctx, annotation.VideoID)
	if err != nil {
		return -1, fmt.Errorf("no video found for the given video id")
	}

	vd, err := time.ParseDuration(video.Duration)
	if err != nil {
		return -1, fmt.Errorf("could not parse duration of the video with id %d", video.ID)
	}

	if vd < et {
		return -1, fmt.Errorf("end time is greater than video duration")
	}

	return u.r.AddAnnotation(ctx, annotation)
}

func (u usecase) GetAnnotation(ctx context.Context, id int) (models.Annotation, error) {
	return u.r.GetAnnotation(ctx, id)
}

func (u usecase) GetAnnotationsByVideoID(ctx context.Context, videoID int) ([]models.Annotation, error) {
	return u.r.GetAnnotationsByVideoID(ctx, videoID)
}

func (u usecase) UpdateAnnotation(ctx context.Context, annotation models.Annotation) error {
	st, err := time.ParseDuration(annotation.StartTime)
	if err != nil {
		return fmt.Errorf("start time is invalid - %s", st)
	}

	et, err := time.ParseDuration(annotation.EndTime)
	if err != nil {
		return fmt.Errorf("end time is invalid - %s", annotation.EndTime)
	}

	if st >= et {
		return fmt.Errorf("start time should be greater than end time")
	}

	video, err := u.GetVideo(ctx, annotation.VideoID)
	if err != nil {
		return fmt.Errorf("no video found for the given video id")
	}

	vd, err := time.ParseDuration(video.Duration)
	if err != nil {
		return fmt.Errorf("could not parse duration of the video with id %d", video.ID)
	}

	if vd < et {
		return fmt.Errorf("end time is greater than video duration")
	}

	return u.r.UpdateAnnotation(ctx, annotation)
}

func (u usecase) DeleteAnnotation(ctx context.Context, id int) error {
	return u.r.DeleteAnnotation(ctx, id)
}
