create table pet_connection(
		id1 VARCHAR(36) NOT NULL DEFAULT 'default' ,
		id2 VARCHAR(36) NOT NULL DEFAULT 'default',
		FOREIGN KEY(id1) REFERENCES pet(id) ON DELETE CASCADE,
		FOREIGN KEY(id2) REFERENCES pet(id)	ON DELETE CASCADE,
		CONSTRAINT PK_pet_connection PRIMARY KEY (id1, id2)
);