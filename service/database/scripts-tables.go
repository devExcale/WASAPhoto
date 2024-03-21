package database

var qCreateTableUser = `
	CREATE TABLE user
	(
		user_uuid    TEXT(36) PRIMARY KEY,
		username     TEXT(20)  NOT NULL UNIQUE,
		display_name TEXT(40)  NOT NULL,
		picture_url  TEXT      NULL,
		ts_created   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`

var qCreateTablePost = `
	CREATE TABLE post
	(
		post_uuid   TEXT(36) PRIMARY KEY,
		author_uuid TEXT(36) NOT NULL,
		caption     TEXT,
		image       BLOB     NOT NULL,
		timestamp   TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);
`

var qCreateTableComment = `
	CREATE TABLE post_comment
	(
		comment_uuid TEXT(36) PRIMARY KEY,
		post_uuid    TEXT(36) NOT NULL,
		author_uuid  TEXT(36) NOT NULL,
		comment      TEXT     NOT NULL,
		timestamp    TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (post_uuid) REFERENCES post (post_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`

var qCreateTableLike = `
	CREATE TABLE post_like
	(
		post_uuid TEXT(36) NOT NULL,
		user_uuid TEXT(36) NOT NULL,
		PRIMARY KEY (post_uuid, user_uuid),
		FOREIGN KEY (post_uuid) REFERENCES post (post_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (user_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`
