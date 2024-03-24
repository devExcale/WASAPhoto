package database

import (
	"errors"
	"github.com/gofrs/uuid"
)

func (db *appdbimpl) GetCommentByUUID(commentUUID string) (Comment, error) {

	var comment Comment

	// Get comment
	var err = db.c.QueryRow(qSelectCommentByUUID, commentUUID).Scan(
		&comment.UUID,
		&comment.Comment,
		&comment.PostUUID,
		&comment.AuthorUUID,
		&comment.CreatedAt,
	)

	return comment, err
}

func (db *appdbimpl) GetCommentsByPost(postUUID string) ([]Comment, error) {

	var comments = make([]Comment, 0)

	// Get comments
	rows, err := db.c.Query(qSelectCommentsByPost, postUUID)
	if err != nil {
		return comments, err
	}

	// Close rows at the end
	defer func() {
		_ = rows.Close()
	}()

	// Map rows to comments
	for rows.Next() {

		var comment Comment
		err = rows.Scan(
			&comment.UUID,
			&comment.Comment,
			&comment.PostUUID,
			&comment.AuthorUUID,
			&comment.CreatedAt,
		)

		if err != nil {
			return []Comment{}, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}

func (db *appdbimpl) SetComment(comment *Comment) error {

	// Check required fields
	if comment.Comment == "" {
		return errors.New("required field comment has not been found")
	}
	if comment.PostUUID == "" {
		return errors.New("required field post has not been found")
	}
	if comment.AuthorUUID == "" {
		return errors.New("required field author has not been found")
	}

	// Create new UUID if not present
	if comment.UUID == "" {
		var genId, _ = uuid.NewV7()
		comment.UUID = genId.String()
	}

	_, err := db.c.Exec(qUpsertComment,
		comment.UUID, comment.Comment, comment.PostUUID, comment.AuthorUUID, // insert values
		comment.Comment, // update values
	)

	return err
}

func (db *appdbimpl) DeleteComment(commentUUID string) error {

	_, err := db.c.Exec(qDeleteComment, commentUUID)

	return err
}
