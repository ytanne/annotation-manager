package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ytanne/annotation-manager/pkg/models"
)

func (u usecase) AddVideo(ctx context.Context, video models.Video) (int, error) {
	_, err := time.ParseDuration(video.Duration)
	if err != nil {
		return -1, fmt.Errorf("could not parse duration of the video with id %d", video.ID)
	}

	return u.r.AddVideo(ctx, video)
}

func (u usecase) GetVideo(ctx context.Context, id int) (models.Video, error) {
	return u.r.GetVideo(ctx, id)
}

func (u usecase) UpdateVideo(ctx context.Context, video models.Video) error {
	_, err := time.ParseDuration(video.Duration)
	if err != nil {
		return fmt.Errorf("could not parse duration of the video with id %d", video.ID)
	}

	return u.r.UpdateVideo(ctx, video)
}

func (u usecase) DeleteVideo(ctx context.Context, id int) error {
	return u.r.DeleteVideo(ctx, id)
}
