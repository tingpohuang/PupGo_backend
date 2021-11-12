create table event(
		id varbinary(16) NOT NULL PRIMARY KEY,
		holder_id varbinary(16) NOT NULL,
		start_date timestamp,
		end_date timestamp,
		image varchar(500),
		limit_user_num int,
		limit_pet_num int,
		description varchar(300)
);