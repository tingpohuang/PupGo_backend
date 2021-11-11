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


Table pet_recommend {
id1 varbinary[16] [pk]
id2 varbinary[16] [pk]
score double
status int
}

Table event {
id varchar[36]
start_date timestamp
end_date timestamp
image varchar[500]
limit_user_num int
limit_dog_num int
description varchar[300]
}

Table event_holder{
id varchar[36]
holder_id varchar[36]
}

Table event_participant{
id varchar[36]
participant_id varchar[36]
}

Table notification{
id varchar[36]
title varchar[50]
description varchar[300]
created_at timestamp
url varchar[500]

type int
user_id varchar[36]
event_id varchar[36]

}

Table user_notification{
id varchar[36]
user_id varchar[36]

}
 
 
 
