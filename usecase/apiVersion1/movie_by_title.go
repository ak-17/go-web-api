package apiVersion1

import (
	"context"

	"github.com/ak-17/go-simple-web-api/model"
)

func (api apiVersion1) GetMovieByTitle(ctx context.Context, title string) (model.Movie, error) {
	return api.repo.GetMovieByTitle(ctx, title)
}
