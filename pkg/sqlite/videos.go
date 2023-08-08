package sqlite

import (
	"context"
	"database/sql"

	"github.com/ytanne/annotation-manager/pkg/errors"
	"github.com/ytanne/annotation-manager/pkg/models"
)

func (c client) AddVideo(ctx context.Context, video models.Video) (int, error) {
	r, err := c.DB.ExecContext(ctx, "INSERT INTO videos (title, description, duration, url) VALUES (?, ?, ?, ?) RETURNING id", video.Title, video.Description, video.Duration, video.URL)
	if err != nil {
		return -1, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (c client) GetVideo(ctx context.Context, id int) (models.Video, error) {
	var video models.Video

	err := c.DB.QueryRowContext(ctx, "SELECT id, title, description, duration, url, created_at, updated_at, deleted_at FROM videos WHERE id = ? and deleted_at is NULL", id).Scan(&video.ID, &video.Title, &video.Description, &video.Duration, &video.URL, &video.CreatedAt, &video.UpdatedAt, &video.DeletedAt)
	if err == sql.ErrNoRows {
		return models.Video{}, errors.ErrNotFound
	} else if err != nil {
		return models.Video{}, err
	}

	return video, nil
}

func (c client) UpdateVideo(ctx context.Context, video models.Video) error {
	_, err := c.DB.ExecContext(ctx, "UPDATE videos SET title = ?, description = ?, duration = ?, url = ?, updated_at = TIME() WHERE id = ? and deleted_at is NULL", video.Title, video.Description, video.Duration, video.URL, video.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c client) DeleteVideo(ctx context.Context, id int) error {
	_, err := c.DB.ExecContext(ctx, "UPDATE videos SET deleted_at = TIME() WHERE id = ? and deleted_at is NULL", id)
	if err != nil {
		return err
	}

	return nil
}
