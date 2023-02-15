package daos

import (
	"context"
	"database/sql/driver"
	"fmt"
	"go-template/models"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/pkg/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Test_createAuthorTx(t *testing.T) {
	type args struct {
		Author models.Author
		ctx    context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Passing user type value",
			args: args{
				Author: models.Author{
					ID:        testutls.MockAuthor().ID,
					FirstName: testutls.MockAuthor().FirstName,
					LastName:  testutls.MockAuthor().FirstName,
				},
				ctx: context.Background(),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
		oldDB := boil.GetDB()
		defer func() {
			boil.SetDB(oldDB)
			db.Close()
		}()
		boil.SetDB(db)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `authors`")).
			WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateAuthor(tt.args.Author, tt.args.ctx)
			fmt.Print(err)
			assert.Equal(t, err, tt.err)

		})
	}
}

func Test_updateAuthorTx(t *testing.T) {
	type args struct {
		Author *models.Author
		ctx    context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Passing user type value",
			args: args{
				Author: testutls.MockAuthor(),
				ctx:    context.Background(),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		err := godotenv.Load("../.env.local")
		if err != nil {
			fmt.Print("error loading .env file")
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		// Inject mock instance into boil.
		oldDB := boil.GetDB()
		defer func() {
			db.Close()
			boil.SetDB(oldDB)
		}()
		boil.SetDB(db)
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `authors` SET `first_name`=?,`last_name`=?," +
			"`updated_at`=?,`deleted_at`=? WHERE `id`=?")).
			WillReturnResult(sqlmock.NewResult(1, 1))
		t.Run(tt.name, func(t *testing.T) {
			_, err := UpdateAuthor(tt.args.Author, tt.args.ctx)
			if err != nil {
				t.Errorf("updateAuthorTx() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}

func TestDeleteAuthor(t *testing.T) {
	type args struct {
		Author models.Author
		ctx    context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Passing user type value",
			args: args{
				Author: models.Author{},
				ctx:    context.Background(),
			},
			err: nil,
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Inject mock instance into boil.
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)
	// delete user
	result := driver.Result(driver.RowsAffected(1))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `authors` WHERE `id`=?")).
		WillReturnResult(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeleteAuthor(tt.args.Author, tt.args.ctx)
			assert.Equal(t, err, tt.err)
		})
	}
}

func TestDeleteAuthorByID(t *testing.T) {
	type args struct {
		authorID int
		ctx      context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Delete pass case",
			args: args{
				authorID: 1,
				ctx:      context.Background(),
			},
		},
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Inject mock instance into boil.
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)
	result := driver.Result(driver.RowsAffected(1))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `authors` WHERE (`authors`.`id` = ?);")).
		WillReturnResult(result)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeleteAuthorByID(tt.args.authorID, tt.args.ctx)
			if err != nil {
				t.Errorf("DeleteAuthorByID() error = %v, wantErr %v", err, tt.err)
				return
			}

		})
	}
}

func TestFetchAllAuthors(t *testing.T) {
	oldDB := boil.GetDB()
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
	query := regexp.QuoteMeta("SELECT `authors`.* FROM `authors`;")
	tests := []struct {
		name      string
		dbQueries []testutls.QueryData
		err       error
	}{
		{
			name: "Failed to find all users with count",
			err:  fmt.Errorf("sql: no rows in sql"),
		},
		{
			name: "Successfully find all users with count",
			err:  nil,
			dbQueries: []testutls.QueryData{
				{
					Query: query,
					DbResponse: sqlmock.NewRows([]string{"id", "email", "token"}).AddRow(
						testutls.MockID,
						testutls.MockEmail,
						testutls.MockToken),
				},
				{
					Query:      regexp.QuoteMeta("SELECT COUNT(*) FROM `users`;"),
					DbResponse: sqlmock.NewRows([]string{"count"}).AddRow(testutls.MockCount),
				},
			},
		},
	}
	for _, tt := range tests {
		if tt.err != nil {
			mock.ExpectQuery(query).
				WithArgs().
				WillReturnError(fmt.Errorf("this is some error"))
		}
		for _, dbQuery := range tt.dbQueries {
			mock.ExpectQuery(dbQuery.Query).
				WithArgs().
				WillReturnRows(dbQuery.DbResponse)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchAllAuthors(context.Background())
			if err != nil {
				assert.Equal(t, true, tt.err != nil)
			} else {
				assert.Equal(t, err, tt.err)
				assert.Equal(t, got[0].FirstName, "")
				assert.Equal(t, got[0].LastName, "")
				assert.Equal(t, got[0].ID, int(testutls.MockID))

			}
		})
	}
	boil.SetDB(oldDB)
	db.Close()
}

func TestFetchAuthorByID(t *testing.T) {
	type args struct {
		ID  int
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Fail on finding user token",
			args: args{
				ID:  testutls.MockAuthor().ID,
				ctx: context.Background(),
			},
			err: errors.WithStack(errors.New("models: unable to select from authors: bind failed to execute query: ")),
		},
		{
			name: "Passing an email",
			args: args{
				ID:  testutls.MockAuthor().ID,
				ctx: context.Background(),
			},
			err: nil,
		},
	}

	query := regexp.QuoteMeta("select * from `authors` where `id`=?")
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Inject mock instance into boil.
	oldDB := boil.GetDB()
	defer func() {
		db.Close()
		boil.SetDB(oldDB)
	}()
	boil.SetDB(db)

	for _, tt := range tests {
		if tt.name == "Fail on finding user token" {
			mock.ExpectQuery(query).
				WithArgs().
				WillReturnError(fmt.Errorf(""))
		} else {
			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
			mock.ExpectQuery(query).
				WithArgs().
				WillReturnRows(rows)
		}

		t.Run(tt.name, func(t *testing.T) {
			_, err := FetchAuthorByID(tt.args.ID, tt.args.ctx)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			}

		})
	}
}
