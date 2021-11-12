create table user_device(
		user_id varbinary(16) NOT NULL,
		device_id varbinary(16) NOT NULL
		PRIMARY KEY ( user_id )
        -- INDEX device_id
);