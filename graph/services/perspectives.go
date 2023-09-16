package services

import (
	"context"
	"review-pull-request-back-end/graph/db"
	"review-pull-request-back-end/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type perspectiveService struct {
	exec boil.ContextExecutor
}

func (u *perspectiveService) CreatePerspective(ctx context.Context, input model.NewPerspective) (*model.Perspective, error) {
	newPerspective := db.Perspective{
		ID:   uuid.New().String(),
		Text: input.Text}
	err := newPerspective.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.Perspective{
		ID:   newPerspective.ID,
		Text: input.Text,
	}, nil
}
