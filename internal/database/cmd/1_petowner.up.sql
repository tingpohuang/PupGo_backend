create table petowner(
		user_id varbinary(16) NOT NULL,
		pet_id varbinary(16) NOT NULL
		PRIMARY KEY ( user_id )
        INDEX pet_id
	 );