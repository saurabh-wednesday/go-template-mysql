package resolver

import (
	"context"
	"fmt"
	"go-template/gqlmodels"
	"go-template/pkg/utl/cnvrttogql"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Test_queryResolver_GetAuthors(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	tests := []struct {
		name string
		fields
		want *gqlmodels.AuthorPayload
		err  error
	}{
		{
			name: "success case fetching users",
			fields: fields{
				Resolver: &Resolver{},
			},
			want: &gqlmodels.AuthorPayload{
				Authors: cnvrttogql.AuthorsToGraphqlAuthors(testutls.MockAuthors()),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		err := godotenv.Load("../.env.local")
		if err != nil {
			fmt.Println("Error loading .env file")
		}
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		// Inject mock instance into boil
		oldDB := boil.GetDB()
		defer func() {
			db.Close()
			boil.SetDB(oldDB)
		}()
		boil.SetDB(db)

		rows := sqlmock.NewRows([]string{"id", "title"}).
			AddRow(testutls.MockID, "Title")
		mock.ExpectQuery(regexp.QuoteMeta("SELECT `authors`.* FROM `authors`;")).
			WithArgs().
			WillReturnRows(rows)
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.fields.Resolver.Query().GetAuthors(context.Background())
			assert.Equal(t, tt.err, err)
		})
	}
}
