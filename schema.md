//// -- LEVEL 1
//// -- Tables and References

// Creating tables
Table users as U {
id varbinary[16] [pk] // auto-increment
cooldown timestamp
created_at timestamp
name varchar[50]
gender int
birthday timestamp

}


Table user_device {
user_id varbinary[16] [pk]
device_id varbinary[16] [pk]
}

Ref: user_device.user_id > U.id

Table petowner {
user_id varbinary[16] [pk] // FK
pet_id varbinary[16] [pk]  // FK
}

Ref: U.id < petowner.user_id
Ref: pet.id < petowner.pet_id

Table pet {
id varbinary[16] [pk]
name varchar[50]
image varchar[500]
gender int
breed varchar[50]
isCastration boolean
birthday timestamp

}

// bi-direction
Table pet_connection {
id1 varbinary[16] [pk]
id2 varbinary[16] [pk]
}

Ref: pet_connection.id1 > pet.id
Ref: pet_connection.id2 > pet.id


// uni-direction
Table pet_recommend {
id1 varbinary[16] [pk]
id2 varbinary[16]
score double
status int
}

Ref: pet_recommend.id2 > pet.id


Table event {
id varbinary[16] [pk]
holder_id varbinary[16]
start_date timestamp
end_date timestamp
image varchar[500]
limit_user_num int
limit_dog_num int
description varchar[300]
}

Ref: event.holder_id > U.id


Table event_participant{
event_id varbinary[16] [pk]
participant_id varbinary[16] [pk]
pet_id varbinary[16] [pk]
status int

}

Ref: event_participant.event_id > event.id
Ref: event_participant.participant_id > U.id
Ref: event_participant.pet_id > pet.id

