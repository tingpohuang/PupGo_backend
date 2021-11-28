create table pet(
		id VARCHAR(36) NOT NULL PRIMARY KEY,
		name varchar(50),
		image varchar(500),
		gender int,
		breed varchar(50),
		isCastration boolean,
		birthday timestamp,
		description varchar(500)
	 );