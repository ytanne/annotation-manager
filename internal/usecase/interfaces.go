package usecase

import (
	"context"

	"github.com/ytanne/annotation-manager/pkg/models"
)

type annotationRepo interface {
	AddAnnotation(context.Context, models.Annotation) (int, error)
	GetAnnotation(context.Context, int) (models.Annotation, error)
	GetAnnotationsByVideoID(context.Context, int) ([]models.Annotation, error)
	UpdateAnnotation(context.Context, models.Annotation) error
	DeleteAnnotation(context.Context, int) error
}

type videoRepo interface {
	AddVideo(context.Context, models.Video) (int, error)
	GetVideo(context.Context, int) (models.Video, error)
	UpdateVideo(context.Context, models.Video) error
	DeleteVideo(context.Context, int) error
}

type repository interface {
	annotationRepo
	videoRepo
}
