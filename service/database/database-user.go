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

	_, err := db.c.Exec(
		qUpsertUser,
		user.UUID, user.Username, user.DisplayName, user.PictureURL, // Insert parameters
		user.Username, user.DisplayName, user.PictureURL, // Update parameters
	)

	return err
}

func (db *appdbimpl) DeleteUser(uuid string) error {

	_, err := db.c.Exec(qDeleteUser, uuid)

	return err
}

func (db *appdbimpl) GetUsersWithUsernameSubstr(substring string, loggedUserUUID string) ([]User, error) {

	var users = make([]User, 0)

	// Get users
	rows, err := db.c.Query(qSelectUsersByUsernameSubstr, substring, loggedUserUUID)
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
