create table event_participant(
		event_id VARCHAR(36) NOT NULL DEFAULT 'default',
		participant_id VARCHAR(36) NOT NULL DEFAULT  'default',
		pet_id VARCHAR(36) DEFAULT 'default',
		status int DEFAULT 0,
		FOREIGN KEY (event_id) REFERENCES event(id) ON DELETE CASCADE,
		FOREIGN KEY (participant_id) REFERENCES users (id) ON DELETE CASCADE,
		FOREIGN KEY (pet_id) REFERENCES pet (id) ON DELETE CASCADE,
		CONSTRAINT PK_event_participant PRIMARY KEY (event_id, participant_id,pet_id)
);
