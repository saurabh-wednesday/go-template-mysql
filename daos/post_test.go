package daos

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go-template/models"
	"go-template/testutls"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestCreatePost(t *testing.T) {
	type args struct {
		post models.Post
		ctx  context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Success case",
			args: args{
				post: models.Post{
					ID: 0,
					AuthorID: null.Int{
						Int:   1,
						Valid: true,
					},
					Title: testutls.MockPost().Title,
					Body:  testutls.MockPost().Title,
				},
				ctx: context.Background(),
			},
			err: nil,
		},
	}
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
	oldDB := boil.GetDB()
	defer func() {
		boil.SetDB(oldDB)
		db.Close()
	}()
	boil.SetDB(db)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `post`")).
		WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := CreatePost(tt.args.post, tt.args.ctx)
			assert.Equal(t, err, tt.err)
		})
	}
}

func Test_createPostTx(t *testing.T) {
	type args struct {
		post models.Post
		ctx  context.Context
		tx   *sql.Tx
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "",
			args: args{
				post: models.Post{
					ID: 0,
					AuthorID: null.Int{
						Int:   1,
						Valid: true,
					},
					Title: "title",
					Body:  "body",
				},
				ctx: nil,
				tx:  &sql.Tx{},
			},
			err: nil,
		},
	}
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
	oldDB := boil.GetDB()
	defer func() {
		boil.SetDB(oldDB)
		db.Close()
	}()
	boil.SetDB(db)
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `post`")).
		WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreatePostTx(tt.args.post, tt.args.ctx, tt.args.tx)
			assert.Equal(t, tt.err, err)

		})
	}
}

func Test_updatePostTx(t *testing.T) {
	type args struct {
		post *models.Post
		ctx  context.Context
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Passing user update action",
			args: args{
				post: &models.Post{
					ID:    1,
					Title: "",
					Body:  "",
				},
				ctx: context.Background(),
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
	mock.ExpectExec(regexp.QuoteMeta("UPDATE `post` SET `author_id`=?,`title`=?," +
		"`body`=?,`updated_at`=?,`deleted_at`=? WHERE `id`=?")).
		WillReturnResult(sqlmock.NewResult(1, 1))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UpdatePost(tt.args.post, tt.args.ctx)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeletePost(t *testing.T) {
	type args struct {
		post models.Post
		ctx  context.Context
	}
	tests := []struct {
		name string
		args args
		want int
		err  error
	}{
		{
			name: "Passing user type value",
			args: args{
				post: models.Post{},
				ctx:  context.Background(),
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
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `post` WHERE `id`=?")).
		WillReturnResult(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeletePost(tt.args.post, tt.args.ctx)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestDeletePostByAuthorID(t *testing.T) {
	type args struct {
		authorID int
		ctx      context.Context
	}
	tests := []struct {
		name string
		args args
		want int
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
	// delete user
	result := driver.Result(driver.RowsAffected(1))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM `post` WHERE (`post`.`author_id` = ?);")).
		WillReturnResult(result)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := DeletePostByAuthorID(tt.args.authorID, tt.args.ctx)
			assert.Equal(t, err, tt.err)

		})
	}
}

func TestFetchAllPosts(t *testing.T) {
	oldDB := boil.GetDB()
	mock, db, _ := testutls.SetupEnvAndDB(t, testutls.Parameters{})
	query := regexp.QuoteMeta("SELECT `post`.* FROM `post`;")
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		args      args
		dbQueries []testutls.QueryData
		err       error
	}{
		{
			name: "Failed to find all users with count",
			args: args{
				ctx: nil,
			},
			dbQueries: []testutls.QueryData{},
			err:       fmt.Errorf("sql: no rows in sql"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
			got, err := FetchAllPosts(tt.args.ctx)

			if err != nil {
				assert.Equal(t, true, tt.err != nil)

			} else {
				assert.Equal(t, err, tt.err)
				assert.Equal(t, got[0].AuthorID, 0)
				assert.Equal(t, got[0].Body, "title")

			}
		})
	}
	boil.SetDB(oldDB)
	db.Close()
}

func TestFetchPostByID(t *testing.T) {
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
			err: errors.WithStack(errors.New("models: unable to select from post: bind failed to execute query: ")),
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

	query := regexp.QuoteMeta("select * from `post` where `id`=?")
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
			_, err := FetchPostByID(tt.args.ID, tt.args.ctx)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			}

		})
	}
}

func TestFetchPostByauthorID(t *testing.T) {
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
				ID:  testutls.MockPost().ID,
				ctx: context.Background(),
			},
			err: errors.WithStack(errors.New("models: failed to assign all query results to Post slice:" +
				" bind failed to execute query: ")),
		},
		{
			name: "Passing an email",
			args: args{
				ID:  testutls.MockPost().ID,
				ctx: context.Background(),
			},
			err: nil,
		},
	}

	query := regexp.QuoteMeta("SELECT `post`.* FROM `post` WHERE (`post`.`id` = ?);")
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
			_, err := FetchPostByauthorID(tt.args.ID, tt.args.ctx)
			if err != nil {
				assert.Equal(t, err.Error(), tt.err.Error())
			}
		})
	}
}
