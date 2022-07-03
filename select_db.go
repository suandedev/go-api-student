package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

type Student struct {
    Id         int
    Email       string
    First_Name string
	Last_Name string
}

func main() {

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_student")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    res, err := db.Query("SELECT * FROM students")

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next() {

        var student Student
        err := res.Scan(&student.Id, &student.Email, &student.First_Name, &student.Last_Name)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v\n", student)
    }
}