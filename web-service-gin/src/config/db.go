package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DB_USERNAME = "root"
const DB_POSWORD = "9900"
const DB_NAME = "user_orm"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}
func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_POSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME

	fmt.Println("dns : ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "t_",
			// table name prefix, table for `User` would be `t_users`
			SingularTable: true,
			// use singular table name, table for `User` would be `user` with this option enabled
			// NoLowerCase: true,
			// skip the snake_casing of names
			// NameReplacer: strings.NewReplacer("CID", "Cid"),
			// use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		// fmt.Println("Error connection to database : error =%v", err)
		return nil
	}
	return db
}
