import (
    "fmt"
    "os"
)


func createDB(name string) {
	DB_RESOURCE_NAME := os.Getenv("DB_RESOURCE_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_URL := os.Getenv("DB_URL")
	DB_PORT := os.Getenv("DB_PORT")
	sql_url := fmt.Sprint(DB_USERNAME, ":", DB_PASSWORD,"@tcp(",DB_URL,":",DB_PORT,")/")
	db, err := sql.Open(DB_RESOURCE_NAME, sql_url)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_,err = db.Exec("CREATE DATABASE "+name)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec("USE "+name)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}
 }