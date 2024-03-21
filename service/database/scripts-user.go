package database

const qSelectUserByUUID = `
	SELECT u.user_uuid,
	       u.username,
	       u.display_name,
	       'TODO',
	       (SELECT COUNT(*) FROM post p WHERE p.author_uuid = u.user_uuid) num_posts,
	       -1 as num_followers,
	       -1 as num_following,
	       u.ts_created
	FROM user u
	WHERE lower(user_uuid) = lower(?)
	LIMIT 1`

const qSelectUserByUsername = `
	SELECT u.user_uuid,
	       u.username,
	       u.display_name,
	       'TODO',
	       (SELECT COUNT(*) FROM post p WHERE p.author_uuid = u.user_uuid) num_posts,
	       -1 as num_followers,
	       -1 as num_following,
	       u.ts_created
	FROM user u
	WHERE username = ?
	LIMIT 1`

const qUpsertUser = `
	INSERT INTO user(user_uuid, username, display_name, picture_url, ts_created)
	VALUES(?, ?, ?, ?, current_timestamp)
	ON CONFLICT (user_uuid) DO
	UPDATE
	SET username = ?, display_name = ?, picture_url = ?`

const qDeleteUser = `
	DELETE FROM user
	WHERE lower(user_uuid) = lower(?)`
