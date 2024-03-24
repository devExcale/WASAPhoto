package database

import "database/sql"

func (db *appdbimpl) UserFollows(followerUUID, followedUUID string) (bool, error) {

	var isfollowed int
	var err = db.c.QueryRow(qSelectFollow, followerUUID, followedUUID).Scan(&isfollowed)

	return isfollowed == 1, err
}

func (db *appdbimpl) GetFollowerUsers(userUUID string) ([]User, error) {

	var slice = make([]User, 0)

	var rows, err = db.c.Query(qSelectFollowers, userUUID)
	if err != nil {
		return []User{}, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// Map rows to users
	for rows.Next() {

		var user = User{}
		err = rows.Scan(
			&user.UUID,
			&user.Username,
			&user.DisplayName,
			&user.PictureURL,
			&user.CreatedAt,
		)

		if err != nil {
			return []User{}, err
		}

		slice = append(slice, user)
	}

	// Create array from slice
	var array = make([]User, len(slice))
	copy(array, slice)

	return array, nil
}

func (db *appdbimpl) GetFollowedUser(userUUID string) ([]User, error) {

	var slice = make([]User, 0)

	var rows, err = db.c.Query(qSelectFollowed, userUUID)
	if err != nil {
		return []User{}, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	// Map rows to users
	for rows.Next() {

		var user = User{}
		err = rows.Scan(
			&user.UUID,
			&user.Username,
			&user.DisplayName,
			&user.PictureURL,
			&user.CreatedAt,
		)

		if err != nil {
			return []User{}, err
		}

		slice = append(slice, user)
	}

	// Create array from slice
	var array = make([]User, len(slice))
	copy(array, slice)

	return array, nil
}

func (db *appdbimpl) AddFollow(followerUUID, followedUUID string) error {

	_, err := db.c.Exec(qInsertFollow, followerUUID, followedUUID)

	return err
}

func (db *appdbimpl) DeleteFollow(followerUUID, followedUUID string) error {

	_, err := db.c.Exec(qDeleteFollow, followerUUID, followedUUID)

	return err
}
