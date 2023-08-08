package usecase

import (
	"context"
	"fmt"

	"github.com/ytanne/annotation-manager/pkg/models"
)

func (u usecase) AddAnnotation(ctx context.Context, annotation models.Annotation) (int, error) {
	if annotation.StartTime >= annotation.EndTime {
		return -1, fmt.Errorf("start time should be smaller than end time. start time - %d, end time - %d", annotation.StartTime, annotation.EndTime)
	}

	video, err := u.GetVideo(ctx, annotation.VideoID)
	if err != nil {
		return -1, fmt.Errorf("no video found for the given video id")
	}

	if video.Duration < annotation.EndTime {
		return -1, fmt.Errorf("end time should be smaller or equal to video duration. video duration - %d, end time - %d", video.Duration, annotation.EndTime)
	}

	id, err := u.r.AddAnnotation(ctx, annotation)
	if err != nil {
		return -1, fmt.Errorf("failed to add annotation for video id - %d", annotation.VideoID)
	}

	return id, nil
}

func (u usecase) GetAnnotation(ctx context.Context, id int) (models.Annotation, error) {
	return u.r.GetAnnotation(ctx, id)
}

func (u usecase) GetAnnotationsByVideoID(ctx context.Context, videoID int) ([]models.Annotation, error) {
	return u.r.GetAnnotationsByVideoID(ctx, videoID)
}

func (u usecase) UpdateAnnotation(ctx context.Context, annotation models.Annotation) error {
	if annotation.StartTime >= annotation.EndTime {
		return fmt.Errorf("start time should be smaller than end time. start time - %d, end time - %d", annotation.StartTime, annotation.EndTime)
	}

	video, err := u.GetVideo(ctx, annotation.VideoID)
	if err != nil {
		return fmt.Errorf("no video found for the given video id")
	}

	if video.Duration < annotation.EndTime {
		return fmt.Errorf("end time should be smaller or equal to video duration. video duration - %d, end time - %d", video.Duration, annotation.EndTime)
	}

	err = u.r.UpdateAnnotation(ctx, annotation)
	if err != nil {
		return fmt.Errorf("failed to update annotation for video id - %d", annotation.VideoID)
	}

	return nil
}

func (u usecase) DeleteAnnotation(ctx context.Context, id int) error {
	return u.r.DeleteAnnotation(ctx, id)
}
