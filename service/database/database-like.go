package database

func (db *appdbimpl) UserHasLikedPost(postUUID, userUUID string) (bool, error) {

	var liked int
	var err = db.c.QueryRow(qSelectUserHasLikedPost, postUUID, userUUID).Scan(&liked)

	return liked == 1, err
}

func (db *appdbimpl) GetUsersLikedPost(postUUID string) ([]User, error) {

	var slice = make([]User, 0)

	var rows, err = db.c.Query(qSelectUsersLikedPost, postUUID)
	if err != nil {
		return []User{}, err
	}

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

	// Check for errors
	err = rows.Err()
	if err != nil {
		return []User{}, err
	}

	// Create array from slice
	var array = make([]User, len(slice))
	copy(array, slice)

	return array, nil
}

func (db *appdbimpl) AddLike(postUUID, userUUID string) error {

	_, err := db.c.Exec(qInsertLike, postUUID, userUUID)

	return err
}

func (db *appdbimpl) DeleteLike(postUUID, userUUID string) error {

	_, err := db.c.Exec(qDeleteLike, postUUID, userUUID)

	return err
}

func (db *appdbimpl) DeletePostLikes(postUUID string) error {

	_, err := db.c.Exec(qDeletePostLikes, postUUID)

	return err
}
