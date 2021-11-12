create table pet(
		id varbinary(16) NOT NULL,
		name varchar(50),
		image varchar(500),
		gender int,
		breed varchar(50),
		isCastration boolean,
		birthday timestamp
		PRIMARY KEY ( id )
	 );