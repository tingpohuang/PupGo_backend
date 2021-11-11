import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)
var DropCmd = map[string] string{
	"users": "DROP TABLE IF EXISTS users;",
	"user_device":"DROP TABLE IF EXISTS user_device;",
	"petowner":"DROP TABLE IF EXISTS petowner;",
	"pet":"DROP TABLE IF EXISTS pet;",
	"event":"DROP TABLE IF EXISTS event;",
	"pet_connection":"DROP TABLE IF EXISTS pet_connection;",
	"event_participant":"DROP TABLE IF EXISTS event_participant;",
	"pet_recommend":"DROP TABLE IF EXISTS pet_recommend;"
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