create table petowner(
		user_id varbinary(16) NOT NULL,
		pet_id varbinary(16) NOT NULL
		PRIMARY KEY ( user_id )
		FOREIGN KEY(pet_id) REFERENCES pet(id)
        -- INDEX pet_id
	 );