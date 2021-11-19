create table event_participant(
		event_id VARCHAR(36) NOT NULL PRIMARY KEY,
		participant_id VARCHAR(36) NOT NULL,
		pet_id VARCHAR(36),
		FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE,
		FOREIGN KEY (participant_id) REFERENCES users (id) ON DELETE CASCADE,
		FOREIGN KEY (pet_id) REFERENCES pet (id) ON DELETE CASCADE
);
