package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"go-template/gqlmodels"
)

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*gqlmodels.UsersPayload, error) {
 return nil,nil

}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }