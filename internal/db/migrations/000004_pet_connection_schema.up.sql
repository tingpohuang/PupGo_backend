create table pet_connection(
		id1 VARCHAR(36) NOT NULL,
		id2 VARCHAR(36) NOT NULL,
		FOREIGN KEY(id1) REFERENCES pet(id) ON DELETE CASCADE,
		FOREIGN KEY(id2) REFERENCES pet(id)	ON DELETE CASCADE,
		CONSTRAINT PK PRIMARY KEY (id1, id2)
);