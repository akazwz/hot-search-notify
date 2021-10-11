package utils

import (
	"database/sql"
	"github.com/akazwz/hot-search-notify/init"
	"time"
)

func getAllSubWords() {
	db, err := sql.Open("mysql", init.CFG.Username+":"+init.CFG.Password+"@/"+init.CFG.DBName)
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

	db.Prepare("SELECT * FROM contents as t WHERE JSON_CONTAINS(t.contents, ?)")
}
