package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/pkg/utl/cnvrttogql"
)

// GetAuthors is the resolver for the getAuthors field.
func (r *queryResolver) GetAuthors(ctx context.Context) (*gqlmodels.AuthorPayload, error) {
	authors, err := daos.FetchAllAuthors(ctx)
	if err != nil {
		return nil, err
	}

	gqlAuthors := cnvrttogql.AuthorsToGraphqlAuthors(authors)
	return &gqlmodels.AuthorPayload{
		Authors: gqlAuthors,
	}, nil
}

// Query returns gqlmodels.QueryResolver implementation.
func (r *Resolver) Query() gqlmodels.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
