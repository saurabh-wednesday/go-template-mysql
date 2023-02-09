package daos

import (
	"context"
	"database/sql"
	"go-template/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreateAuthor creates a new Author based on the input provided to the Author model
func CreateAuthor(Author models.Author, ctx context.Context) (models.Author, error) {
	return createAuthorTx(Author, ctx, nil)
}

// createAuthorTx holds a transcation happening within the db to create a Author
func createAuthorTx(Author models.Author, ctx context.Context, tx *sql.Tx) (models.Author, error) {
	contextExecutor := getContextExecutor(tx)
	err := Author.Insert(ctx, contextExecutor, boil.Infer())
	return Author, err
}

// UpdateAuthor updates a given Author values
func UpdateAuthor(Author *models.Author, ctx context.Context) (models.Author, error) {
	return updateAuthorTx(Author, ctx, nil)
}

// updateAuthorTx creates a transaction for the updation of Author
func updateAuthorTx(Author *models.Author, ctx context.Context, tx *sql.Tx) (models.Author, error) {
	contextExecutor := getContextExecutor(tx)

	_, err := Author.Update(ctx, contextExecutor, boil.Infer())
	return *Author, err
}

// DeleteAuthor deletes a given Author
func DeleteAuthor(Author models.Author, ctx context.Context) (int, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := Author.Delete(ctx, contextExecutor)
	return int(rowsAffected), err
}

// DeleteAuthor deletes a given Author
func DeleteAuthorByID(authorID int, ctx context.Context) (int, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := models.Authors(models.AuthorWhere.ID.EQ(authorID)).DeleteAll(ctx, contextExecutor)
	return int(rowsAffected), err
}

// FetchAllAuthors returns all the Authors created
func FetchAllAuthors(ctx context.Context) (models.AuthorSlice, error) {
	contextExecutor := getContextExecutor(nil)
	Authors, err := models.Authors().All(ctx, contextExecutor)
	if err != nil {
		return models.AuthorSlice{}, err
	}
	return Authors, err
}

// FetchAuthorByID returns the Author associated with the id
func FetchAuthorByID(ID int, ctx context.Context) (*models.Author, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindAuthor(ctx, contextExecutor, ID)
}
