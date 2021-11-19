create table petowner(
		user_id varbinary(16) NOT NULL PRIMARY KEY,
		pet_id varbinary(16) NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(pet_id) REFERENCES pet(id) ON DELETE CASCADE
	 );