create table user_location(
		user_id VARCHAR(36) NOT NULL PRIMARY KEY,
		latitude     FLOAT (8,5) ,
		longitude    FLOAT(8,5) ,
		country varchar(50),
		state varchar(50),
		city varchar(50),
		address varchar (50),
		FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
