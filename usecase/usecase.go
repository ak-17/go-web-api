package usecase

import (
	"context"

	"github.com/ak-17/go-simple-web-api/model"
)

type Usecase interface {
	GetMovieByTitle(ctx context.Context, name string) (model.Movie, error)
}
