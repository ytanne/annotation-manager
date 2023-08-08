package controller

import (
	"context"

	"github.com/ytanne/annotation-manager/pkg/models"
)

type VideoUsecase interface {
	AddVideo(ctx context.Context, video models.Video) (int, error)
	GetVideo(ctx context.Context, id int) (models.Video, error)
	UpdateVideo(ctx context.Context, video models.Video) error
	DeleteVideo(ctx context.Context, id int) error
}

type AnnotationUsecase interface {
	AddAnnotation(ctx context.Context, annotation models.Annotation) (int, error)
	GetAnnotation(ctx context.Context, id int) (models.Annotation, error)
	GetAnnotationsByVideoID(ctx context.Context, videoID int) ([]models.Annotation, error)
	UpdateAnnotation(ctx context.Context, annotation models.Annotation) error
	DeleteAnnotation(ctx context.Context, id int) error
}

type Usecase interface {
	VideoUsecase
	AnnotationUsecase
}
