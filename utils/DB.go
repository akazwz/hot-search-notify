package utils

import (
	"database/sql"
	"fmt"
	"github.com/akazwz/hot-search-notify/inital"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func GetAllSubWords() {
	db, err := sql.Open("mysql", inital.CFG.Username+":"+inital.CFG.Password+"@mysql(47.96.24.50:3306)/"+inital.CFG.DBName)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10)

	stmtOut, err := db.Prepare("SELECT * FROM contents as t WHERE JSON_CONTAINS(t.contents, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmtOut)

	var rows string
	err = stmtOut.QueryRow("赵文卓").Scan(&rows)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(rows)
}
