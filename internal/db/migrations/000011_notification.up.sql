create table notification(
		notification_id VARCHAR(36) NOT NULL PRIMARY KEY,
		notification_type int,
        user_id VARCHAR(36),
        pet_id VARCHAR(36),
        event_id VARCHAR(36),
        created_at timestamp,
        payload VARCHAR(1024),
        has_read boolean,
);
