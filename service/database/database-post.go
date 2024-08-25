package database

import (
	"errors"
	"github.com/gofrs/uuid"
)

func (db *appdbimpl) GetPost(postUUID string) (Post, error) {

	var post = Post{}

	// Get post
	err := db.c.QueryRow(qSelectPost, postUUID).Scan(
		&post.UUID,
		&post.AuthorUUID,
		&post.Caption,
		&post.ImageURL,
		&post.NComments,
		&post.NLikes,
		&post.CreatedAt,
	)

	return post, err
}

func (db *appdbimpl) GetPostWithLike(postUUID, userUUID string) (Post, error) {

	var post = Post{}

	// Get post
	err := db.c.QueryRow(qSelectPostWithLike, postUUID, userUUID).Scan(
		&post.UUID,
		&post.AuthorUUID,
		&post.AuthorUsername,
		&post.AuthorDisplayName,
		&post.Caption,
		&post.ImageURL,
		&post.NComments,
		&post.NLikes,
		&post.CreatedAt,
		&post.LoggedUserLiked,
	)

	return post, err
}

func (db *appdbimpl) GetPostsByUser(userUUID string) ([]Post, error) {

	var posts = make([]Post, 0)

	// Get posts
	rows, err := db.c.Query(qSelectPostsByUser, userUUID)
	if err != nil {
		return posts, err
	}

	// Map rows to posts
	for rows.Next() {

		var post Post
		err = rows.Scan(
			&post.UUID,
			&post.AuthorUUID,
			&post.AuthorUsername,
			&post.AuthorDisplayName,
			&post.Caption,
			&post.ImageURL,
			&post.NComments,
			&post.NLikes,
			&post.CreatedAt,
		)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	// Check for errors
	err = rows.Err()
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}

func (db *appdbimpl) GetPostsByUserWithLikes(authorUUID, loggedUUID string) ([]Post, error) {

	var posts = make([]Post, 0)

	// Get posts
	rows, err := db.c.Query(qSelectPostsByUserWithLikes, authorUUID, loggedUUID)
	if err != nil {
		return posts, err
	}

	// Map rows to posts
	for rows.Next() {

		var post Post
		err = rows.Scan(
			&post.UUID,
			&post.AuthorUUID,
			&post.AuthorUsername,
			&post.AuthorDisplayName,
			&post.Caption,
			&post.ImageURL,
			&post.NComments,
			&post.NLikes,
			&post.CreatedAt,
			&post.LoggedUserLiked,
		)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	// Check for errors
	err = rows.Err()
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}

func (db *appdbimpl) GetPostsByFollowed(userUUID string) ([]Post, error) {

	var posts = make([]Post, 0)

	// Get posts
	rows, err := db.c.Query(qSelectPostsByFollowed, userUUID)
	if err != nil {
		return posts, err
	}

	// Map rows to posts
	for rows.Next() {

		var post Post
		err = rows.Scan(
			&post.UUID,
			&post.AuthorUUID,
			&post.AuthorUsername,
			&post.AuthorDisplayName,
			&post.Caption,
			&post.ImageURL,
			&post.NComments,
			&post.NLikes,
			&post.CreatedAt,
			&post.LoggedUserLiked,
		)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	// Check for errors
	err = rows.Err()
	if err != nil {
		return []Post{}, err
	}

	return posts, nil
}

func (db *appdbimpl) GetImage(postUUID string) ([]byte, error) {

	var image []byte

	// Get post
	err := db.c.QueryRow(qSelectImage, postUUID).Scan(&image)

	return image, err
}

func (db *appdbimpl) SetPost(post *Post, image []byte) error {

	// Check required fields
	if post.AuthorUUID == "" {
		return errors.New("required field author_uuid has not been found")
	}
	if image == nil {
		return errors.New("required field image has not been found")
	}

	// Create new UUID if not present
	if post.UUID == "" {
		var genId, _ = uuid.NewV7()
		post.UUID = genId.String()
	}

	_, err := db.c.Exec(qUpsertPost,
		post.UUID, post.AuthorUUID, post.Caption, image, // insert values
		post.Caption, // update values
	)

	return err
}

func (db *appdbimpl) DeletePost(uuid string) error {

	_, err := db.c.Exec(qDeletePost, uuid)

	return err
}
