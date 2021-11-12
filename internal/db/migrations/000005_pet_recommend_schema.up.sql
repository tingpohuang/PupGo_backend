create table pet_recommend(
		id1 varbinary(16) NOT NULL PRIMARY KEY,
		id2 varbinary(16) NOT NULL,
		score float(55),
		status int
		FOREIGN KEY(id2) REFERENCES pet(id)
	);