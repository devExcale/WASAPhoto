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
)
