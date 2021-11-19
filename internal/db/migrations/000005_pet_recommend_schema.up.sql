create table pet_recommend(
		id1 VARCHAR(36) NOT NULL,
		id2 VARCHAR(36) NOT NULL,
		score float(20),
		status int,
		FOREIGN KEY(id1) REFERENCES pet(id) ON DELETE CASCADE,
		FOREIGN KEY(id2) REFERENCES pet(id)	ON DELETE CASCADE,
		CONSTRAINT PK_pet_recommend PRIMARY KEY (id1, id2)
	);