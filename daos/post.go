package daos

import (
	"context"
	"database/sql"
	"go-template/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// CreatePost creates a new post based on the input provided to the post model
func CreatePost(post models.Post, ctx context.Context) (models.Post, error) {
	return createPostTx(post, ctx, nil)
}

// createPostTx holds a transcation happening within the db to create a post
func createPostTx(post models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
	contextExecutor := getContextExecutor(tx)
	err := post.Insert(ctx, contextExecutor, boil.Infer())
	return post, err
}

// UpdatePost updates a given post values
func UpdatePost(post *models.Post, ctx context.Context) (models.Post, error) {
	return updatePostTx(post, ctx, nil)
}

// updatePostTx creates a transaction for the updation of post
func updatePostTx(post *models.Post, ctx context.Context, tx *sql.Tx) (models.Post, error) {
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

// FetchAllPosts returns all the posts created
func FetchAllPosts(ctx context.Context) (models.PostSlice, error) {
	contextExecutor := getContextExecutor(nil)
	posts, err := models.Posts().All(ctx, contextExecutor)
	if err != nil {
		return models.PostSlice{}, err
	}
	return posts, err
}
func FetchPostByID(ID int, ctx context.Context) (*models.Post,error) {
    contextExecutor:=getContextExecutor(nil)
    return models.FindPost(ctx,contextExecutor,ID)
}
