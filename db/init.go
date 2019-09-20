package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/firedial/midas-go/config"
)

func Init() *sql.DB {
    dsn := config.DB_USER + ":" + config.DB_PASSWORD + "@/midas?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }

    return db
}
