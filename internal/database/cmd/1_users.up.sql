create table users(
		id varbinary(16) NOT NULL,
		cooldown timestamp,
		created_at timetstamp NOT NULL,
		name VARCHAR(50),
		gender INT,
		birthday timetstamp
		PRIMARY KEY ( id )
);