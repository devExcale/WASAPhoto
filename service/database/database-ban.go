package database

func (db *appdbimpl) GetBan(issuerUUID, bannedUUID string) (bool, error) {

	var isBanned int
	var err = db.c.QueryRow(qSelectBan, issuerUUID, bannedUUID).Scan(&isBanned)

	return isBanned == 1, err
}

func (db *appdbimpl) AddBan(issuerUUID, bannedUUID string) error {

	_, err := db.c.Exec(qInsertBan, issuerUUID, bannedUUID)

	return err
}

func (db *appdbimpl) DeleteBan(issuerUUID, bannedUUID string) error {

	_, err := db.c.Exec(qDeleteBan, issuerUUID, bannedUUID)

	return err
}
