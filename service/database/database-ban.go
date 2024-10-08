package database

func (db *appdbimpl) IsBanned(bannedUUID, issuerUUID string) (bool, error) {

	var ban int
	var err = db.c.QueryRow(qSelectIsBanned, issuerUUID, bannedUUID).Scan(&ban)

	return ban == 1, err
}

func (db *appdbimpl) GetBannedUsers(issuerUUID string) ([]User, error) {

	var slice = make([]User, 0)

	var rows, err = db.c.Query(qSelectBanned, issuerUUID)
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

func (db *appdbimpl) AddBan(issuerUUID, bannedUUID string) error {

	_, err := db.c.Exec(qInsertBan, issuerUUID, bannedUUID)

	return err
}

func (db *appdbimpl) DeleteBan(issuerUUID, bannedUUID string) error {

	_, err := db.c.Exec(qDeleteBan, issuerUUID, bannedUUID)

	return err
}
