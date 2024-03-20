/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB.
type AppDatabase interface {

	// GetUser retrieves the user with the given UUID.
	GetUser(uuid string) (User, error)

	// SetUser adds or updates a user. No need to provide the UUID for new users.
	// The object passed as parameter will be updated with the inserted data.
	SetUser(user *User) error

	// DeleteUser removes the user with the given UUID.
	DeleteUser(uuid string) error

	// GetPost retrieves the post with the given UUID.
	GetPost(uuid string) (Post, error)

	// SetPost adds or updates a post. No need to provide the UUID for new posts.
	// The object passed as parameter will be updated with the inserted data.
	SetPost(post *Post) error

	// DeletePost removes the post with the given UUID.
	DeletePost(uuid string) error

	//GetComment(uuid string) (Comment, error)
	//SetComment(comment *Comment) error
	//AddLikePost(postUUID, userUUID string) error
	//RemoveLikePost(postUUID, userUUID string) error

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var err error

	// user table
	err = createTableIfNotExists(db, `user`, `
		CREATE TABLE user
		(
			uuid         TEXT(36) PRIMARY KEY,
			username     TEXT(20) NOT NULL UNIQUE,
			display_name TEXT(40) NOT NULL,
			picture      TEXT     NULL
		);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// post table
	err = createTableIfNotExists(db, `post`, `
		CREATE TABLE post
		(
			uuid      TEXT(36) PRIMARY KEY,
			caption   TEXT,
			image     BLOB     NOT NULL,
			timestamp TEXT     NOT NULL DEFAULT current_timestamp,
			user_uuid TEXT(36) NOT NULL,
			FOREIGN KEY (user_uuid) REFERENCES user (uuid)
				ON DELETE CASCADE
				ON UPDATE CASCADE
		);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// post_comment table
	err = createTableIfNotExists(db, `post_comment`, `
		CREATE TABLE post_comment
		(
			uuid      TEXT(36) PRIMARY KEY,
			comment   TEXT     NOT NULL,
			timestamp TEXT     NOT NULL DEFAULT current_timestamp,
			post_uuid TEXT(36) NOT NULL,
			user_uuid TEXT(36) NOT NULL,
			FOREIGN KEY (post_uuid) REFERENCES post (uuid)
				ON UPDATE CASCADE
				ON DELETE CASCADE,
			FOREIGN KEY (user_uuid) REFERENCES user (uuid)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	// post_like table
	err = createTableIfNotExists(db, `post_like`, `
		CREATE TABLE post_like
		(
			post_uuid TEXT(36) NOT NULL,
			user_uuid TEXT(36) NOT NULL,
			PRIMARY KEY (post_uuid, user_uuid),
			FOREIGN KEY (post_uuid) REFERENCES post (uuid)
				ON UPDATE CASCADE
				ON DELETE CASCADE,
			FOREIGN KEY (user_uuid) REFERENCES user (uuid)
				ON UPDATE CASCADE
				ON DELETE CASCADE
		);`)
	if err != nil {
		return nil, fmt.Errorf("error creating database structure: %w", err)
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func createTableIfNotExists(conn *sql.DB, tablename, createStmt string) error {

	if conn == nil {
		return errors.New("missing database connection")
	}

	err := conn.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name=?", tablename).Scan(&tablename)

	if errors.Is(err, sql.ErrNoRows) {
		_, err = conn.Exec(createStmt)
	}

	return err
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
