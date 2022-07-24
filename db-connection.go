package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                              // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("the error has occured in connection", err)
		panic("connection failed" + err.Error())
	} else {
		fmt.Println("The database has been Connected Succesfully", *DB)
		DB.AutoMigrate(&Employee{})
	}
	return DB
}
