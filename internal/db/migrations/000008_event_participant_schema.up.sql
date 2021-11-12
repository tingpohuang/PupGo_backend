create table event_participant(
		event_id1 varbinary(16) NOT NULL PRIMARY KEY,
		participant_id varbinary(16) NOT NULL,
		pet_id varbinary(16)
        -- INDEX participant_id
);
-- CREATE INDEX "event_participant" ("participant_id");
