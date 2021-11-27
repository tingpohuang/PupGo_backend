create table user_notification(
		notification_id VARCHAR(36) NOT NULL PRIMARY KEY,
		notification_type int,
        user_id VARCHAR(36) NOT NULL,
        pet_id VARCHAR(36),
        event_id VARCHAR(36),
        created_at timestamp,
        payload VARCHAR(1024),
        has_read boolean,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (pet_id) REFERENCES pet(id)
);

