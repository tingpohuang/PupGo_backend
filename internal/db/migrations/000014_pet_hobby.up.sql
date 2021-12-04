create table pet_hobby(
                          pet_id VARCHAR(36) NOT NULL DEFAULT 'default' ,
                          hobby_id int NOT NULL DEFAULT '1',
                          FOREIGN KEY(pet_id) REFERENCES pet(id) ON DELETE CASCADE,
                          FOREIGN KEY(hobby_id) REFERENCES hobby(id)	ON DELETE CASCADE,
                          CONSTRAINT PK_pet_hobby PRIMARY KEY (pet_id, hobby_id)
);