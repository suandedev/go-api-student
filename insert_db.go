package main

import (
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

type Student struct {
	Id         int
    Email       string
    First_Name string
	Last_Name string
}

func main(){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_student")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

	sql := "INSERT INTO students(email, first_name, last_name) VALUES ('admin@gmail.com', 'admin','admin')"
			
    res, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

    lastId, err := res.LastInsertId()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The last inserted row id: %d\n", lastId)
}

