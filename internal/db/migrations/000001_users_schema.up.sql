create table users(
		id VARCHAR(36) NOT NULL PRIMARY KEY,
		cooldown timestamp,
		email VARCHAR(100),
		created_at timestamp NOT NULL,
		name VARCHAR(50),
		gender INT,
		birthday timestamp
);