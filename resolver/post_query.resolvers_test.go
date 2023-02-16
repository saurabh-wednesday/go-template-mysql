package resolver

import (
	"context"
	"go-template/daos"
	"go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/cnvrttogql"
	"go-template/testutls"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func Test_queryResolver_GetPosts(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gqlmodels.PostPayload
		err    error
	}{
		{
			name: "success case fetching posts",
			fields: fields{
				Resolver: &Resolver{},
			},
			want: &gqlmodels.PostPayload{
				Posts: cnvrttogql.PostsToGraphQlPosts(testutls.MockPosts()),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &queryResolver{
				Resolver: tt.fields.Resolver,
			}
			gomonkey.ApplyFunc(daos.FetchAllPosts, func(ctx context.Context) (models.PostSlice, error) {
				return testutls.MockPosts(), nil
			})
			_, err := r.GetPosts(tt.args.ctx)
			assert.Equal(t, err, tt.err)

		})
	}
}
