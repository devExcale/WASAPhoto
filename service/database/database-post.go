package database

import (
	"errors"
	"github.com/gofrs/uuid"
)

const qSelectPost = `
	SELECT post_uuid, author_uuid, caption, ts_created
	FROM post
	WHERE lower(post_uuid) = lower(?)`

const qUpsertPost = `
	INSERT INTO post(post_uuid, author_uuid, caption, image, ts_created)
	VALUES(lower(?), ?, ?, ?, current_timestamp)
	ON CONFLICT (post_uuid) DO
	UPDATE SET caption = ?`

const qDeletePost = `
	DELETE FROM post
	WHERE lower(post_uuid) = lower(?)`

func (db *appdbimpl) GetPost(uuid string) (Post, error) {

	var post = Post{}

	err := db.c.QueryRow(qSelectPost, uuid).Scan(&post.UUID, &post.Caption, &post.CreatedAt, &post.AuthorUUID)

	return post, err
}

func (db *appdbimpl) SetPost(post *Post) error {

	// Check required fields
	if post.Caption == "" {
		return errors.New("required field caption has not been found")
	}
	if post.AuthorUUID == "" {
		return errors.New("required field author_uuid has not been found")
	}

	// TODO: image

	// Create new UUID if not present
	if post.UUID == "" {
		var genId, _ = uuid.NewV7()
		post.UUID = genId.String()
	}

	_, err := db.c.Exec(qUpsertPost, post.UUID, post.AuthorUUID, post.Caption, nil)

	return err
}

func (db *appdbimpl) DeletePost(uuid string) error {

	_, err := db.c.Exec(qDeletePost, uuid)

	return err
}
