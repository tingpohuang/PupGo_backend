//// -- LEVEL 1
//// -- Tables and References

// Creating tables
Table users as U {
id varchar[36] [pk, increment] // auto-increment
cooldown timestamp
created_at timestamp

}

Table petowner {
id varchar[36]
pet_list varchar[180]
}


Table pet {
id varchar[36]
pet_owner_id varchar[36]

}

Table user_profile {
id varchar[36]
name varchar[50]
gender int
birthday timestamp

}

Table pet_profile {
id varchar[36]
name varchar[50]
image varchar[500]
gender int
breed varchar[50]
isCastration boolean
birthday timestamp

}

Table pet_connection {
id1 varchar[36]
id2 varchar[36]
}

Table pet_recommend {
id1 varchar[36]
id2 varchar[36]
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



// Creating references
// You can also define relaionship separately
// > many-to-one; < one-to-many; - one-to-one
//Ref: U.country_code > countries.code  
//Ref: merchants.country_code > countries.code

//----------------------------------------------//

//// -- LEVEL 2
//// -- Adding column settings


Table order_items {
order_id int [ref: > orders.id] // inline relationship (many-to-one)
product_id int
quantity int [default: 1] // default value
}

Ref: order_items.product_id > products.id

Table orders {
id int [pk] // primary key
user_id int [not null, unique]
status varchar
created_at varchar [note: 'When order created'] // add column note
}

//----------------------------------------------//

//// -- Level 3
//// -- Enum, Indexes

// Enum for 'products' table below
Enum products_status {
out_of_stock
in_stock
running_low [note: 'less than 20'] // add column note
}

// Indexes: You can define a single or multi-column index
Table products {
id int [pk]
name varchar
merchant_id int [not null]
price int
status products_status
created_at datetime [default: `now()`]

Indexes {
(merchant_id, status) [name:'product_status']
id [unique]
}
}

Table merchants {
id int
country_code int
merchant_name varchar

"created at" varchar
admin_id int [ref: > U.id]
Indexes {
(id, country_code) [pk]
}
}

Table merchant_periods {
id int [pk]
merchant_id int
country_code int
start_date datetime
end_date datetime
}

Ref: products.merchant_id > merchants.id // many-to-one
//composite foreign key
Ref: merchant_periods.(merchant_id, country_code) > merchants.(id, country_code)
