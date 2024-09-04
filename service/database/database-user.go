package database

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
)

func (db *appdbimpl) GetUserFull(userUUID, loggedUserUUID string) (User, error) {

	var user = User{}

	err := db.c.QueryRow(qSelectUserFullByUUID, userUUID, loggedUserUUID).Scan(
		&user.UUID,
		&user.Username,
		&user.DisplayName,
		&user.PictureURL,
		&user.NPosts,
		&user.NFollowers,
		&user.NFollowing,
		&user.CreatedAt,
		&user.LoggedUserFollowed,
		&user.LoggerUserRestricted,
	)

	return user, err
}

func (db *appdbimpl) GetUserBasicByUUID(uuid string) (User, error) {

	var user = User{}

	err := db.c.QueryRow(qSelectUserBasicByUUID, uuid).Scan(
		&user.UUID,
		&user.Username,
		&user.DisplayName,
		&user.PictureURL,
		&user.CreatedAt,
	)

	return user, err
}

func (db *appdbimpl) GetUserBasicByUsername(username string) (User, error) {

	var user = User{}

	err := db.c.QueryRow(qSelectUserBasicByUsername, username).Scan(
		&user.UUID,
		&user.Username,
		&user.DisplayName,
		&user.PictureURL,
		&user.CreatedAt,
	)

	return user, err
}

func (db *appdbimpl) SetUser(user *User) error {

	// Check required fields
	if user.Username == "" {
		return errors.New("required field username has not been found")
	}

	// Set display name to username if not present
	if user.DisplayName == "" {
		user.DisplayName = user.Username
	}

	// Create new UUID if not present
	if user.UUID == "" {
		var genId, _ = uuid.NewV7()
		user.UUID = genId.String()
	}

	_, err := db.c.Exec(
		qUpsertUser,
		user.UUID, user.Username, user.DisplayName, user.PictureURL, // Insert parameters
		user.Username, user.DisplayName, user.PictureURL, // Update parameters
	)

	return err
}

func (db *appdbimpl) DeleteUser(uuid string, tx *sql.Tx) error {

	var err error

	// Init transaction if null
	if tx == nil {

		tx, err = db.c.Begin()
		if err == nil {

			// Commit or rollback new transaction
			defer func() {
				if err == nil {
					err = tx.Commit()
				} else {
					_ = tx.Rollback()
				}
			}()

		}
	}

	// Select user posts
	var posts []Post
	if err == nil {
		posts, err = db.GetPostsByUser(uuid)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
	}

	// Delete all posts
	if err == nil {
		for _, post := range posts {
			err = db.DeletePost(post.UUID, tx)
			if err != nil {
				break
			}
		}
	}

	// Delete following and follower links
	if err == nil {
		_, err = tx.Exec(qDeleteUserFollowerds, uuid)
	}

	// Delete likes
	if err == nil {
		_, err = tx.Exec(qDeleteLikesByUser, uuid)
	}

	// Delete comments
	if err == nil {
		_, err = tx.Exec(qDeleteCommentsByUser, uuid)
	}

	// Delete bans
	if err == nil {
		_, err = tx.Exec(qDeleteBansByOnUser, uuid)
	}

	// Delete user
	if err == nil {
		_, err = tx.Exec(qDeleteUser, uuid)
	}

	return err
}

func (db *appdbimpl) GetUsersSubstringLike(substring string, loggedUserUUID string) ([]User, error) {

	var users = make([]User, 0)

	// Get users
	rows, err := db.c.Query(qSelectUsersSubstringLike, substring, loggedUserUUID)
	if err != nil {
		return users, err
	}

	// Map rows to users
	for rows.Next() {

		var user User
		err = rows.Scan(
			&user.UUID,
			&user.Username,
			&user.DisplayName,
			&user.PictureURL,
			&user.CreatedAt,
		)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	// Check for errors
	err = rows.Err()
	if err != nil {
		return []User{}, err
	}

	return users, nil
}

func (db *appdbimpl) IsUsernameAvailable(username string) (bool, error) {

	var taken bool
	var err = db.c.QueryRow(qIsUsernameTaken, username).Scan(&taken)

	return !taken, err
}
