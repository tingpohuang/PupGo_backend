create table petowner(
		user_id varbinary(16) NOT NULL PRIMARY KEY,
		pet_id varbinary(16) NOT NULL,
		FOREIGN KEY(pet_id) REFERENCES pet(id)
        -- INDEX pet_id
	 );