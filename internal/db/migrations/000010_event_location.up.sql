create table event_location(
		event_id VARCHAR(36) NOT NULL PRIMARY KEY,
		position point,
		country varchar(50),
		state varchar(50),
		city varchar(50),
		FOREIGN KEY (event_id) REFERENCES event (id) ON DELETE CASCADE

);
