package repository

import (
	"context"

	"github.com/ak-17/go-simple-web-api/model"
)

type Repository interface {
	GetMovieByTitle(ctx context.Context, name string) (model.Movie, error)
}
