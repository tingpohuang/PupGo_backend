create table users(
		id varbinary(16) NOT NULL,
		cooldown timestamp,
		created_at timestamp NOT NULL,
		name VARCHAR(50),
		gender INT,
		birthday timestamp
		PRIMARY KEY ( id )
);