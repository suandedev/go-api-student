package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
type Student struct {
    Id         int
    Email       string
    First_Name string
	Last_Name string
}
 
func main() {
	db, e := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_student")
    ErrorCheck(e)
 
    // close database after all work is done
    defer db.Close()
 
    PingDB(db)
 
    //Update db
    stmt, e := db.Prepare("update students set First_Name=? where id=?")
    ErrorCheck(e)
 
    // execute
    res, e := stmt.Exec("Ramesh", "1")
    ErrorCheck(e)
 
    a, e := res.RowsAffected()
    ErrorCheck(e)
 
    fmt.Println(a)
 
    
}
 
func ErrorCheck(err error) {
    if err != nil {
        panic(err.Error())
    }
}
 
func PingDB(db *sql.DB) {
    err := db.Ping()
    ErrorCheck(err)
}