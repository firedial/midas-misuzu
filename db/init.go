package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/firedial/midas-misuzu/config"
)

func Init() *sql.DB {
    dsn := config.DB_USER + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ")/" + config.DB_DATABASE + "?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err.Error())
    }

    return db
}
