create table event(
		id VARCHAR(36) NOT NULL PRIMARY KEY,
		holder_id VARCHAR(36) NOT NULL,
		start_date timestamp,
		end_date timestamp,
		image varchar(500),
		limit_user_num int,
		limit_pet_num int,
		description varchar(300),
		type int,
		FOREIGN KEY (holder_id) REFERENCES pet(id) ON DELETE CASCADE
);