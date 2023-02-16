package daos

import (
	"context"
	"database/sql"
	"go-template/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreatePost creates a new post based on the input provided to the post model
func CreatePost(post models.Post, ctx context.Context) (models.Post, error) {
	return CreatePostTx(post, ctx, nil)
}

// CreatePostTx holds a transcation happening within the db to create a post
func CreatePostTx(post models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
	contextExecutor := getContextExecutor(tx)
	err := post.Insert(ctx, contextExecutor, boil.Infer())
	return post, err
}

// UpdatePost updates a given post values
func UpdatePost(post *models.Post, ctx context.Context) (models.Post, error) {
	return UpdatePostTx(post, ctx, nil)
}

// UpdatePostTx creates a transaction for the updation of post
func UpdatePostTx(post *models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
	contextExecutor := getContextExecutor(tx)
	_, err := post.Update(ctx, contextExecutor, boil.Infer())
	return *post, err
}

// DeletePost deletes a given post
func DeletePost(post models.Post, ctx context.Context) (int, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := post.Delete(ctx, contextExecutor)
	return int(rowsAffected), err
}

// DeletePostByAuthorID deletes all the post by a given author
func DeletePostByAuthorID(authorID int, ctx context.Context) (int, error) {
	contextExecutor := getContextExecutor(nil)
	rowsAffected, err := models.Posts(models.PostWhere.AuthorID.EQ(null.NewInt(authorID, true))).DeleteAll(ctx, contextExecutor)
	if err != nil {
		return -1, err
	}
	return int(rowsAffected), err
}

// FetchAllPosts returns all the posts created
func FetchAllPosts(ctx context.Context) (models.PostSlice, error) {
	contextExecutor := getContextExecutor(nil)
	posts, err := models.Posts().All(ctx, contextExecutor)
	if err != nil {
		return models.PostSlice{}, err
	}
	return posts, err
}

// FetchPostByID returns the post associated with the id
func FetchPostByID(ID int, ctx context.Context) (*models.Post, error) {
	contextExecutor := getContextExecutor(nil)
	return models.FindPost(ctx, contextExecutor, ID)
}

// FetchPostByID returns the post associated with the id
func FetchPostByAuthorID(ID int, ctx context.Context) (*models.PostSlice, error) {
	contextExecutor := getContextExecutor(nil)
	posts, err := models.Posts(models.PostWhere.ID.EQ(ID)).All(ctx, contextExecutor)
	if err != nil {
		return nil, err
	}
	return &posts, nil

}
