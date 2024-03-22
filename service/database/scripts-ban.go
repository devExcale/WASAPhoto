package database

const qSelectBan = `
	SELECT COUNT(*)
	FROM (
		SELECT 1
		FROM user_banned
		WHERE lower(issuer_uuid) = lower(?)
			AND lower(banned_uuid) = lower(?)
	)`

const qInsertBan = `
	INSERT INTO user_banned(issuer_uuid, banned_uuid)
	VALUES(?, ?)
	ON CONFLICT DO NOTHING`

const qDeleteBan = `
	DELETE FROM user
	WHERE lower(user_uuid) = lower(?)`
