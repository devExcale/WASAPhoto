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
