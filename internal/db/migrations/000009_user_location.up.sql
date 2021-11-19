create table user_location(
		user_id varbinary(16) NOT NULL PRIMARY KEY,
		position point,
		country varchar(50),
		state varchar(50),
		city varchar(50),
		FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
