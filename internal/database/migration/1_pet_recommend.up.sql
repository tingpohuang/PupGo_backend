create table pet_recommend(
		id1 varbinary(16) NOT NULL,
		id2 varbinary(16) NOT NULL,
		score float(55),
		status int
		PRIMARY KEY ( id1 )
        INDEX id2
	);