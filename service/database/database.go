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

	// GetUserFull retrieves all user information with the given UUID or Username.
	GetUserFull(param string, filterBy int) (User, error)

	// GetUserBasic retrieves minimal user information with the given UUID.
	GetUserBasic(userUUID string) (User, error)

	// SetUser adds or updates a user. No need to provide the UUID for new users.
	// The object passed as parameter will be updated with the inserted data.
	SetUser(user *User) error

	// DeleteUser removes the user with the given UUID.
	DeleteUser(userUUID string) error

	// GetPost retrieves the post with the given UUID.
	GetPost(postUUID string) (Post, error)

	// GetPostsByUser retrieves all posts by the given user.
	GetPostsByUser(userUUID string) ([]Post, error)

	// GetPostsByFollowed retrieves all posts by the users followed by the given user.
	GetPostsByFollowed(userUUID string) ([]Post, error)

	// GetImage retrieves the image of the post with the given UUID.
	GetImage(postUUID string) ([]byte, error)

	// SetPost adds or updates a post. No need to provide the UUID for new posts.
	// The object passed as parameter will be updated with the inserted data.
	SetPost(post *Post, image []byte) error

	// DeletePost removes the post with the given UUID.
	DeletePost(postUUID string) error

	// IsBanned checks if a user is banned by another user.
	IsBanned(issuerUUID, bannedUUID string) (bool, error)

	// GetBannedUsers retrieves all users banned by the given user.
	GetBannedUsers(issuerUUID string) ([]User, error)

	// AddBan bans a user from seeing another user's content.
	AddBan(issuerUUID, bannedUUID string) error

	// DeleteBan lifts a ban.
	DeleteBan(issuerUUID, bannedUUID string) error

	// UserFollows checks if a user follows another user.
	UserFollows(followerUUID, followedUUID string) (bool, error)

	// GetFollowerUsers retrieves all users that follow the given user.
	GetFollowerUsers(userUUID string) ([]User, error)

	// GetFollowedUsers retrieves all users followed by the given user.
	GetFollowedUsers(userUUID string) ([]User, error)

	// AddFollow adds a follow relationship between two users.
	AddFollow(followerUUID, followedUUID string) error

	// DeleteFollow removes a follow relationship between two users.
	DeleteFollow(followerUUID, followedUUID string) error

	// GetCommentByUUID retrieves a comment by its UUID
	GetCommentByUUID(commentUUID string) (Comment, error)

	// GetCommentsByPost retrieves all comments for a given post
	GetCommentsByPost(postUUID string) ([]Comment, error)

	// SetComment adds a comment to a post
	SetComment(comment *Comment) error

	// DeleteComment removes a comment from a post
	DeleteComment(commentUUID string) error

	// UserHasLikedPost checks if a user has liked a post
	UserHasLikedPost(postUUID, userUUID string) (bool, error)

	// GetUsersLikedPost retrieves all users who liked a post
	GetUsersLikedPost(postUUID string) ([]User, error)

	// AddLike adds a like to a post
	AddLike(postUUID, userUUID string) error

	// DeleteLike removes a like from a post
	DeleteLike(postUUID, userUUID string) error

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

	var err error

	// Check if table exists. If not, the database is empty, and we need to create the structure
	for tablename, createStmt := range tables {
		err = createTableIfNotExists(db, tablename, createStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating table %s: %w", tablename, err)
		}
	}

	return &appdbimpl{c: db}, nil
}

func createTableIfNotExists(conn *sql.DB, tablename, createStmt string) error {

	if conn == nil {
		return errors.New("missing database connection")
	}

	const qCheckTable = `SELECT name FROM sqlite_master WHERE type IN ('table', 'view') AND name=?`

	err := conn.QueryRow(qCheckTable, tablename).Scan(&tablename)

	if errors.Is(err, sql.ErrNoRows) {
		_, err = conn.Exec(createStmt)
	}

	return err
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
