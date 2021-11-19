create table petowner(
		user_id VARCHAR(36) NOT NULL PRIMARY KEY,
		pet_id VARCHAR(36) NOT NULL,
		FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY(pet_id) REFERENCES pet(id) ON DELETE CASCADE
	 );