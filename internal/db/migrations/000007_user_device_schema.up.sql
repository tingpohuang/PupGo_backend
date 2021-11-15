create table user_device(
		user_id varbinary(16) NOT NULL PRIMARY KEY,
		device_id varbinary(16) NOT NULL
        -- INDEX device_id
);