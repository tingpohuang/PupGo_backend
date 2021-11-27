create table user_token(
		user_id VARCHAR(36) NOT NULL PRIMARY KEY,
		token VARCHAR (500),
		FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
