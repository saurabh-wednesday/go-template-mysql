package resolver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-template/daos"
	"go-template/gqlmodels"
	fm "go-template/gqlmodels"
	"go-template/models"
	"go-template/pkg/utl/convert"
	"go-template/testutls"
	"strconv"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func Test_mutationResolver_CreatePost(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input gqlmodels.PostCreateInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gqlmodels.Post
		err    error
	}{
		{
			name: "fail case on post creation",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx:   context.Background(),
				input: fm.PostCreateInput{},
			},
			err: errors.New("invalid input"),
		},
		{
			name: "Success",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: context.Background(),
				input: fm.PostCreateInput{
					Title:    testutls.MockPost().Title,
					Body:     testutls.MockPost().Body,
					AuthorID: "1",
				},
			},
			want: &fm.Post{
				ID:        fmt.Sprint(testutls.MockPost().ID),
				AuthorID:  "1",
				Title:     testutls.MockPost().Title,
				Body:      &testutls.MockPost().Body,
				CreatedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
				UpdatedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
				DeletedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			postPatch := gomonkey.ApplyFunc(
				daos.CreatePostTx, func(post models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
					return *testutls.MockPost(), nil
				})
			authorFetchPatch := gomonkey.ApplyFunc(daos.FetchAuthorByID, func(ID int, ctx context.Context) (*models.Author, error) {
				return testutls.MockAuthor(), nil
			})
			strConvAtoi := gomonkey.ApplyFunc(strconv.Atoi, func(s string) (int, error) {
				return 1, nil
			})
			defer authorFetchPatch.Reset()
			defer strConvAtoi.Reset()
			defer postPatch.Reset()
			_, err := r.CreatePost(tt.args.ctx, tt.args.input)
			if err != nil {
				assert.Equal(t, err, tt.err)
			}
		})
	}
}

func Test_mutationResolver_UpdatePost(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input gqlmodels.PostUpdateInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gqlmodels.Post
		err    error
	}{

		{
			name: "Success",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: context.Background(),
				input: gqlmodels.PostUpdateInput{

					ID:    fmt.Sprint(testutls.MockPost().ID),
					Title: &testutls.MockPost().Title,
					Body:  &testutls.MockPost().Body,
				},
			},
			want: &fm.Post{
				ID:        fmt.Sprint(testutls.MockPost().ID),
				AuthorID:  "1",
				Title:     testutls.MockPost().Title,
				Body:      &testutls.MockPost().Body,
				CreatedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
				UpdatedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
				DeletedAt: convert.NullDotTimeToPointerInt(testutls.MockPost().CreatedAt),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			postPatch := gomonkey.ApplyFunc(
				daos.UpdatePost, func(post *models.Post, ctx context.Context) (models.Post, error) {
					return *testutls.MockPost(), nil
				})
			strConvAtoi := gomonkey.ApplyFunc(strconv.Atoi, func(s string) (int, error) {
				return 0, nil
			})
			PostFetchPatch := gomonkey.ApplyFunc(daos.FetchPostByID, func(ID int, ctx context.Context) (*models.Post, error) {
				return testutls.MockPost(), nil
			})
			defer strConvAtoi.Reset()
			defer postPatch.Reset()
			defer PostFetchPatch.Reset()
			_, err := r.UpdatePost(tt.args.ctx, tt.args.input)
			assert.Equal(t, err, tt.err)

		})
	}
}

func Test_mutationResolver_DeletePost(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input gqlmodels.PostDeleteInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gqlmodels.PostDeleteResponse
		err    error
	}{
		{
			name: "success case",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: nil,
				input: fm.PostDeleteInput{
					ID: strconv.Itoa(testutls.MockAuthor().ID),
				},
			},
			want: &fm.PostDeleteResponse{
				RowsAffected: 1,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			postPatch := gomonkey.ApplyFunc(
				daos.DeletePost, func(post models.Post, ctx context.Context) (int, error) {
					return 1, nil
				})
			fetchPostByIdPatch := gomonkey.ApplyFunc(
				daos.FetchPostByID, func(ID int, ctx context.Context) (*models.Post, error) {
					return testutls.MockPost(), nil
				})

			defer postPatch.Reset()
			defer fetchPostByIdPatch.Reset()
			_, err := r.DeletePost(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.err, err)

		})
	}
}
