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
 
    // delete data
    stmt, e := db.Prepare("delete from students where id=?")
    ErrorCheck(e)
 
    // delete 1st student
    res, e := stmt.Exec("1")
    ErrorCheck(e)
 
    // affected rows
    a, e := res.RowsAffected()
    ErrorCheck(e)
 
    fmt.Println(a) // 1
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