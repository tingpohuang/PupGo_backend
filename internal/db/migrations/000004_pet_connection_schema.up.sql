create table pet_connection(
		id1 varbinary(16) NOT NULL PRIMARY KEY,
		id2 varbinary(16) NOT NULL,
		FOREIGN KEY(id2) REFERENCES pet(id)
);