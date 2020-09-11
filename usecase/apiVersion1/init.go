package apiVersion1

import (
	"github.com/ak-17/go-simple-web-api/repository"
	"github.com/ak-17/go-simple-web-api/usecase"
)

func New(repo repository.Repository) (usecase.Usecase, error) {
	api := apiVersion1{repo: repo}
	return api, nil
}
