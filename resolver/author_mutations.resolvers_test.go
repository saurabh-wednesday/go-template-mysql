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

func Test_mutationResolver_CreateAuthor(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input fm.AuthorCreateInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *fm.Author
		err    error
	}{
		{
			name: "fail case on user creation",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx:   context.Background(),
				input: fm.AuthorCreateInput{},
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
				input: fm.AuthorCreateInput{
					FirstName: testutls.MockUser().FirstName.String,
					LastName:  testutls.MockUser().LastName.String,
				},
			},
			want: &fm.Author{
				ID:        fmt.Sprint(testutls.MockUser().ID),
				FirstName: testutls.MockUser().FirstName.String,
				LastName:  testutls.MockUser().LastName.String,
				CreatedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().CreatedAt),
				DeletedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().DeletedAt),
				UpdatedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().UpdatedAt),
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			authorPatch := gomonkey.ApplyFunc(
				daos.CreateAuthorTx, func(Author models.Author, ctx context.Context, tx *sql.Tx) (models.Author, error) {
					return *testutls.MockAuthor(), nil
				})
			defer authorPatch.Reset()
			_, err := r.CreateAuthor(tt.args.ctx, tt.args.input)
			assert.Equal(t, err, tt.err)
		})
	}
}

func Test_mutationResolver_UpdateAuthor(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input gqlmodels.AuthorUpdateInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *fm.Author
		err    error
	}{
		{
			name: "fail case on user update",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: context.Background(),
				input: fm.AuthorUpdateInput{
					ID:        strconv.Itoa(testutls.MockAuthor().ID),
					FirstName: &testutls.MockAuthor().FirstName,
					LastName:  &testutls.MockAuthor().FirstName,
				},
			},
			want: &fm.Author{
				ID:        fmt.Sprint(testutls.MockUser().ID),
				FirstName: testutls.MockUser().FirstName.String,
				LastName:  testutls.MockUser().LastName.String,
				CreatedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().CreatedAt),
				DeletedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().DeletedAt),
				UpdatedAt: convert.NullDotTimeToPointerInt(testutls.MockAuthor().UpdatedAt),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			authorPatch := gomonkey.ApplyFunc(
				daos.UpdateAuthorTx, func(Author *models.Author, ctx context.Context, tx *sql.Tx) (models.Author, error) {
					return *testutls.MockAuthor(), nil
				})
			findauthorPatch := gomonkey.ApplyFunc(
				daos.FetchAuthorByID, func(ID int, ctx context.Context) (*models.Author, error) {
					return testutls.MockAuthor(), nil
				})
			defer authorPatch.Reset()
			defer findauthorPatch.Reset()
			_, err := r.UpdateAuthor(tt.args.ctx, tt.args.input)
			assert.Equal(t, err, tt.err)
		})
	}
}

func Test_mutationResolver_DeleteAuthor(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input gqlmodels.AuthorDeleteInput
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *gqlmodels.AuthorDeleteResponse
		err    error
	}{
		{
			name: "success case",
			fields: fields{
				Resolver: &Resolver{},
			},
			args: args{
				ctx: nil,
				input: fm.AuthorDeleteInput{
					ID: strconv.Itoa(testutls.MockAuthor().ID),
				},
			},
			want: &fm.AuthorDeleteResponse{
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

			authorPatch := gomonkey.ApplyFunc(
				daos.DeletePostByAuthorID, func(authorID int, ctx context.Context) (int, error) {
					return 1, nil
				})
			findauthorPatch := gomonkey.ApplyFunc(
				daos.DeleteAuthorByID, func(authorID int, ctx context.Context) (int, error) {
					return 1, nil
				})
			defer authorPatch.Reset()
			defer findauthorPatch.Reset()
			_, err := r.DeleteAuthor(tt.args.ctx, tt.args.input)
			assert.Equal(t, tt.err, err)

		})
	}
}
