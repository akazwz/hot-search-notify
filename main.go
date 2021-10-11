package main

import (
	"database/sql"
	"fmt"
	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/internal/sub"
)

func main() {
	fmt.Println("hello, notify")
	inital.VP = inital.InitViper()
	if inital.VP == nil {
		fmt.Println("配置文件初始化失败")
	}

	inital.GDB = inital.InitDB()
	if inital.GDB != nil {
		inital.CreateTables(inital.GDB)
		db, _ := inital.GDB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
			}
		}(db)
	} else {
		fmt.Println("数据库连接失败")
		return
	}

	sub.NotifySub()
}
