create table pet_connection(
		id1 varbinary(16) NOT NULL,
		id2 varbinary(16) NOT NULL
		PRIMARY KEY ( id1 )
        INDEX id2
);