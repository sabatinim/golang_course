package db_init

import (
	"github.com/jinzhu/gorm"
	//"github.com/go-sql-driver/mysql"
	//"github.com/mattn/go-sqlite3"
	"fmt"
	_ "github.com/lib/pq"
)

var myDb gorm.DB

func init() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=book_library sslmode=disable")
	fmt.Println("Err", err)
	db.LogMode(true)
	db.DB()
	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.CreateTable(&User{})
	db.CreateTable(&Address{})
	db.CreateTable(&Email{})
	db.CreateTable(&Language{})
	myDb = db
}

func GetDb() gorm.DB {
	return myDb
}
