package database

import (
	"errors"
	"github.com/gofrs/uuid"
)

const qSelectUser = `
	SELECT user_uuid, username, display_name, ts_created
	FROM user
	WHERE lower(user_uuid) = lower(?)`

const qUpsertUser = `
	INSERT INTO user(user_uuid, username, display_name, picture_url, ts_created)
	VALUES(?, ?, ?, ?, current_timestamp)
	ON CONFLICT (user_uuid) DO
	UPDATE
	SET username = ?, display_name = ?, picture_url = ?`

const qDeleteUser = `
	DELETE FROM user
	WHERE lower(user_uuid) = lower(?)`

func (db *appdbimpl) GetUser(uuid string) (User, error) {

	var user = User{}

	err := db.c.QueryRow(qSelectUser, uuid).Scan(&user.UUID, &user.Username, &user.DisplayName, &user.CreatedAt)

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

	_, err := db.c.Exec(qUpsertUser, user.UUID, user.Username, user.DisplayName, user.Username, user.DisplayName)

	return err
}

func (db *appdbimpl) DeleteUser(uuid string) error {

	_, err := db.c.Exec(qDeleteUser, uuid)

	return err
}
