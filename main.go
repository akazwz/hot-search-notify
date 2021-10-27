package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/internal/sub"
	"github.com/robfig/cron/v3"
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

	//generatePDF(fmt.Sprintf("%s", time.Now().Format("2006-01-02-15-04-05")))
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("时区加载失败")
	}

	// 开启定时任务
	c := cron.New(cron.WithLocation(location))
	_, err = c.AddFunc("* * * * * ", func() {
		log.Println("")
		sub.NotifySub()
	})

	if err != nil {
		log.Fatal("定时任务添加失败", err)
	}
	c.Run()
	c.Start()
}
