create table notification(
		notification_id VARCHAR(36) NOT NULL PRIMARY KEY,
		notification_type int,
        user_id VARCHAR(36) NOT NULL,
        created_at timestamp,
        payload VARCHAR(1024) NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
