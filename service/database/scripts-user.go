package database

const qSelectUserFullByUUID = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.num_posts,
		u.num_followed,
		u.num_following,
		u.ts_created
	FROM
		user_full u
	WHERE
		lower(u.user_uuid) = lower(?)
	LIMIT 1
`

const qSelectUserFullByUsername = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.num_posts,
		u.num_followed,
		u.num_following,
		u.ts_created
	FROM
		user_full u
	WHERE
		u.username = ?
	LIMIT 1
`

const qSelectUserBasicByUUID = `
	SELECT
		user_uuid,
		username,
		display_name,
		picture_url,
		ts_created
	FROM
		user
	WHERE
		lower(user_uuid) = lower(?)
	LIMIT 1
`

const qUpsertUser = `
	INSERT INTO user(user_uuid, username, display_name, picture_url, ts_created)
	VALUES(?, ?, ?, ?, current_timestamp)
	ON CONFLICT (user_uuid) DO
	UPDATE
	SET username = ?, display_name = ?, picture_url = ?`

const qDeleteUser = `
	DELETE FROM user
	WHERE lower(user_uuid) = lower(?)`
