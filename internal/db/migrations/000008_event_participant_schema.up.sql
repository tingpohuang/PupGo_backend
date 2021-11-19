create table event_participant(
		event_id varbinary(16) NOT NULL PRIMARY KEY,
		participant_id varbinary(16) NOT NULL,
		pet_id varbinary(16),
		FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE,
		FOREIGN KEY (participant_id) REFERENCES users (id) ON DELETE CASCADE,
		FOREIGN KEY (pet_id) REFERENCES pet (id) ON DELETE CASCADE
);
