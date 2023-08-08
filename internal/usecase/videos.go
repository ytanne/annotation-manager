package usecase

import (
	"context"
	"fmt"

	"github.com/ytanne/annotation-manager/pkg/models"
)

func (u usecase) AddVideo(ctx context.Context, video models.Video) (int, error) {
	if video.Duration <= 0 {
		return -1, fmt.Errorf("duration of the video with id %d shouldn't negative or zero", video.ID)
	}

	return u.r.AddVideo(ctx, video)
}

func (u usecase) GetVideo(ctx context.Context, id int) (models.Video, error) {
	return u.r.GetVideo(ctx, id)
}

func (u usecase) UpdateVideo(ctx context.Context, video models.Video) error {
	if video.Duration <= 0 {
		return fmt.Errorf("duration of the video with id %d shouldn't negative or zero", video.ID)
	}

	return u.r.UpdateVideo(ctx, video)
}

func (u usecase) DeleteVideo(ctx context.Context, id int) error {
	return u.r.DeleteVideo(ctx, id)
}
