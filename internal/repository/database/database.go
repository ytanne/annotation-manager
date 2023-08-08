package database

import (
	"context"

	"github.com/ytanne/annotation-manager/internal/repository"
	"github.com/ytanne/annotation-manager/pkg/models"
)

type repo struct {
	db repository.Database
}

func NewRepository(db repository.Database) repo {
	return repo{
		db: db,
	}
}

func (r repo) AddVideo(ctx context.Context, video models.Video) (int, error) {
	return r.db.AddVideo(ctx, video)
}

func (r repo) GetVideo(ctx context.Context, id int) (models.Video, error) {
	return r.db.GetVideo(ctx, id)
}

func (r repo) UpdateVideo(ctx context.Context, video models.Video) error {
	return r.db.UpdateVideo(ctx, video)
}

func (r repo) DeleteVideo(ctx context.Context, id int) error {
	return r.db.DeleteVideo(ctx, id)
}

func (r repo) AddAnnotation(ctx context.Context, annotation models.Annotation) (int, error) {
	return r.db.AddAnnotation(ctx, annotation)
}

func (r repo) GetAnnotation(ctx context.Context, id int) (models.Annotation, error) {
	return r.db.GetAnnotation(ctx, id)
}

func (r repo) GetAnnotationsByVideoID(ctx context.Context, videoID int) ([]models.Annotation, error) {
	return r.db.GetAnnotationsByVideoID(ctx, videoID)
}

func (r repo) UpdateAnnotation(ctx context.Context, annotation models.Annotation) error {
	return r.db.UpdateAnnotation(ctx, annotation)
}

func (r repo) DeleteAnnotation(ctx context.Context, id int) error {
	return r.db.DeleteAnnotation(ctx, id)
}
