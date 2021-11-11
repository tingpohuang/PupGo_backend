import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
var CreateCmd = map[string] string{
	"users": "create table users(
		id varbinary(16) NOT NULL,
		cooldown timestamp,
		created_at timetstamp NOT NULL,
		name VARCHAR(50),
		gender INT,
		birthday timetstamp
		PRIMARY KEY ( id )
	 );",
	 "user_device":"create table user_device(
		user_id varbinary(16),
		device_id varbinary(16)
		PRIMARY KEY ( user_id )
	 );",
	 "petowner":"create table petowner(
		user_id varbinary(16),
		pet_id varbinary(16)
		PRIMARY KEY ( user_id )
	 );",
	 "pet":"create table pet(
		id varbinary(16),
		name varchar(50),
		image varchar(500),
		gender int,
		breed varchar(50),
		isCastration boolean,
		brithday timetstamp
		PRIMARY KEY ( id )
	 );",
	 "event":"create table event(
		id varbinary(16),
		holder_id varbinary(16),
		start_date timestamp,
		end_date timestamp,
		image varchar(500),
		limit_user_num int,
		limit_pet_num int,
		description varchar(300)
		PRIMARY KEY ( id )
	 );",
	 "pet_connection":"create table pet_connection(
		id1 varbinary(16),
		id2 varbinary(16)
		PRIMARY KEY ( id1 )
	 );",
	 "event_participant":"create table event_participant(
		event_id1 varbinary(16),
		participant_id varbinary(16),
		pet_id varbinary(16)
		PRIMARY KEY ( event_id1 )
	 );",
	"pet_recommend":"create table pet_recommend(
		id1 varbinary(16),
		id2 varbinary(16),
		score double,
		status int
		PRIMARY KEY ( id1 )
	);"
}
func Main() {
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        panic(err)
    } else if err = db.Ping(); err != nil {
        panic(err)
    }
    defer db.Close()
}

func CreateTables(){
	res, err := db.Exec("INSERT INTO mytable (some_text) VALUES (?)", "hello world")
    if err != nil {
        panic(err)
    }
}