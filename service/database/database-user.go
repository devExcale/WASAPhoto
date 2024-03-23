package database

import (
	"errors"
	"github.com/gofrs/uuid"
)

const (
	FilterByUUID = iota
	FilterByUsername
)

func (db *appdbimpl) GetUserFull(param string, filterBy int) (User, error) {

	var user = User{}
	var query string

	switch filterBy {
	case FilterByUUID:
		query = qSelectUserFullByUUID
	case FilterByUsername:
		query = qSelectUserFullByUsername
	default:
		return user, errors.New("invalid filterBy value")
	}

	err := db.c.QueryRow(query, param).Scan(
		&user.UUID,
		&user.Username,
		&user.DisplayName,
		&user.PictureURL,
		&user.NPosts,
		&user.NFollowers,
		&user.NFollowing,
		&user.CreatedAt)

	return user, err
}

func (db *appdbimpl) GetUserBasic(uuid string) (User, error) {

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

	_, err := db.c.Exec(qUpsertUser,
		user.UUID, user.Username, user.DisplayName, "", // Insert parameters
		user.Username, user.DisplayName, "") // Update parameters

	return err
}

func (db *appdbimpl) DeleteUser(uuid string) error {

	_, err := db.c.Exec(qDeleteUser, uuid)

	return err
}
