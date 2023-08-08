package sqlite

import (
	"context"

	"github.com/ytanne/annotation-manager/pkg/models"
)

func (c client) AddAnnotation(ctx context.Context, annotation models.Annotation) (int, error) {
	r, err := c.DB.ExecContext(ctx, "INSERT INTO annotations (video_id, start_time, end_time, text) VALUES (?, ?, ?, ?) RETURNING id", annotation.VideoID, annotation.StartTime, annotation.EndTime, annotation.Text)
	if err != nil {
		return -1, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (c client) GetAnnotation(ctx context.Context, id int) (models.Annotation, error) {
	var annotation models.Annotation

	err := c.DB.QueryRowContext(ctx, "SELECT id, video_id, start_time, end_time, text, created_at, updated_at, deleted_at FROM annotations WHERE id = ?", id).Scan(&annotation.ID, &annotation.VideoID, &annotation.StartTime, &annotation.EndTime, &annotation.Text, &annotation.CreatedAt, &annotation.UpdatedAt, &annotation.DeletedAt)
	if err != nil {
		return models.Annotation{}, err
	}

	return annotation, nil
}

func (c client) GetAnnotationsByVideoID(ctx context.Context, videoID int) ([]models.Annotation, error) {
	var annotations []models.Annotation

	rows, err := c.DB.QueryContext(ctx, "SELECT id, video_id, start_time, end_time, text, created_at, updated_at, deleted_at FROM annotations WHERE video_id = ?", videoID)
	if err != nil {
		return []models.Annotation{}, err
	}

	for rows.Next() {
		var annotation models.Annotation

		err := rows.Scan(&annotation.ID, &annotation.VideoID, &annotation.StartTime, &annotation.EndTime, &annotation.Text, &annotation.CreatedAt, &annotation.UpdatedAt, &annotation.DeletedAt)
		if err != nil {
			return []models.Annotation{}, err
		}

		annotations = append(annotations, annotation)
	}

	return annotations, nil
}

func (c client) UpdateAnnotation(ctx context.Context, annotation models.Annotation) error {
	_, err := c.DB.ExecContext(ctx, "UPDATE annotations SET start_time = ?, end_time = ?, text = ?, updated_at = TIME() WHERE id = ?", annotation.StartTime, annotation.EndTime, annotation.Text, annotation.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c client) DeleteAnnotation(ctx context.Context, id int) error {
	_, err := c.DB.ExecContext(ctx, "UPDATE annotations SET deleted_at = TIME() WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
