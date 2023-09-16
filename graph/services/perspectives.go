package services

import (
	"context"
	"review-pull-request-back-end/graph/db"
	"review-pull-request-back-end/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type perspectiveService struct {
	exec boil.ContextExecutor
}

func (p *perspectiveService) CreatePerspective(ctx context.Context, input model.NewPerspective) (*model.Perspective, error) {
	newPerspective := db.Perspective{Text: input.Text}
	err := newPerspective.Insert(ctx, p.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.Perspective{
		ID:   newPerspective.ID,
		Text: input.Text,
	}, nil
}

func (p *perspectiveService) QueryPerspectives(ctx context.Context) ([]*model.Perspective, error) {
	dbPerspectives, err := db.Perspectives(
		qm.Select(
			db.PerspectiveColumns.ID,
			db.PerspectiveColumns.Text,
		)).All(ctx, p.exec)
	if err != nil {
		return nil, err
	}

	var perspectives []*model.Perspective
	for _, dbPerspective := range dbPerspectives {
		link := &model.Perspective{
			ID:   dbPerspective.ID,
			Text: dbPerspective.Text,
		}
		perspectives = append(perspectives, link)
	}

	return perspectives, nil
}
