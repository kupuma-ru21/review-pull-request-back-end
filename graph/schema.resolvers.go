package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"review-pull-request-back-end/graph/model"
)

// CreatePerspective is the resolver for the createPerspective field.
func (r *mutationResolver) CreatePerspective(ctx context.Context, input model.NewPerspective) (*model.Perspective, error) {
	return r.Srv.CreatePerspective(ctx, input)
}

// Perspectives is the resolver for the perspectives field.
func (r *queryResolver) Perspectives(ctx context.Context) ([]*model.Perspective, error) {
	var perspectives []*model.Perspective
	dummyPerspective := model.Perspective{
		Text: "text",
	}
	perspectives = append(perspectives, &dummyPerspective)
	return perspectives, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
