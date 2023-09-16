package services

import (
	"context"
	"review-pull-request-back-end/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PerspectiveService interface {
	CreatePerspective(ctx context.Context, input model.NewPerspective) (*model.Perspective, error)
}

type Services interface {
	PerspectiveService
}

type services struct {
	*perspectiveService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		perspectiveService: &perspectiveService{exec: exec},
	}
}
