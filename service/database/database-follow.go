package database

func (db *appdbimpl) GetFollow(followerUUID, followedUUID string) (bool, error) {

	var isfollowed int
	var err = db.c.QueryRow(qSelectFollow, followerUUID, followedUUID).Scan(&isfollowed)

	return isfollowed == 1, err
}

func (db *appdbimpl) AddFollow(followerUUID, followedUUID string) error {

	_, err := db.c.Exec(qInsertFollow, followerUUID, followedUUID)

	return err
}

func (db *appdbimpl) DeleteFollow(followerUUID, followedUUID string) error {

	_, err := db.c.Exec(qDeleteFollow, followerUUID, followedUUID)

	return err
}
