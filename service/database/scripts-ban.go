package database

const qSelectIsBanned = `
	SELECT
		COUNT(*)
	FROM (
		SELECT 1
		FROM user_banned
		WHERE lower(issuer_uuid) = lower(?)
		  AND lower(banned_uuid) = lower(?)
	)`

const qSelectBanned = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.ts_created
	FROM
		user_banned b
	JOIN
		user u ON u.user_uuid = b.banned_uuid
	WHERE lower(b.issuer_uuid) = lower(?)
`

const qInsertBan = `
	INSERT INTO user_banned(issuer_uuid, banned_uuid)
	VALUES(?, ?)
	ON CONFLICT DO NOTHING
`

const qDeleteBan = `
	DELETE FROM user_banned
	WHERE lower(issuer_uuid) = lower(?)
	  AND lower(banned_uuid) = lower(?)
`

const qDeleteBansByOnUser = `
	DELETE FROM user_banned
	WHERE lower(banned_uuid) = lower(?1)
	 OR lower(issuer_uuid) = lower(?1)
`
