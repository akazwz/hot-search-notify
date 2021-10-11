package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Res struct {
	Code int64     `json:"code"`
	Msg  string    `json:"msg"`
	Data HotSearch `json:"data"`
}

type HotSearch struct {
	Time     string            `json:"time"`
	Searches []SingleHotSearch `json:"searches"`
}

type SingleHotSearch struct {
	Rank    int    `json:"rank"`
	Content string `json:"content"`
	Link    string `json:"link"`
	Hot     int    `json:"hot"`
	Tag     string `json:"tag"`
	Icon    string `json:"icon"`
}

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

	utils.GetAllSubWords()

	// 请求接口,获取当前热搜
	response, err := http.Get("https://hs.hellozwz.com/hot-searches/current")
	if err != nil {
		log.Println("请求失败")
		return
	}
	if response.StatusCode != http.StatusOK {
		log.Println("请求失败")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read error")
		return
	}
	// json 结构到 res
	res := Res{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("json error")
		return
	}
	// 热搜数组
	searches := res.Data.Searches
	// 储存热搜内容数组
	contentsArr := make([]string, 0)
	for i := 0; i < len(searches); i++ {
		content := searches[i].Content
		contentsArr = append(contentsArr, content)
	}
	// 全部订阅词汇数组
	subWordsArr := []string{
		"王嘉尔",
		"女",
		"假的",
	}
	// 热搜包含的订阅词汇
	filterWordsArr := getFilterWords(subWordsArr, contentsArr)
	// 符合订阅的热搜内容
	filterContentArr := make([]string, 0)
	for i := 0; i < len(contentsArr); i++ {
		for j := 0; j < len(filterWordsArr); j++ {
			isContains := strings.Contains(contentsArr[i], filterWordsArr[j])
			if isContains {
				filterContentArr = append(filterContentArr, contentsArr[i])
			}
		}
	}
	log.Println("订阅", filterWordsArr)
	log.Println("内容", filterContentArr)
}

// 获取热搜包含的订阅词汇
func getFilterWords(subWordsArr, contentsArr []string) []string {
	filterArr := make([]string, 0)
	for j := 0; j < len(subWordsArr); j++ {
		filterStr := regexp.MustCompile(subWordsArr[j]).FindString(strings.Join(contentsArr, "|"))
		if len(filterStr) > 0 {
			filterArr = append(filterArr, filterStr)
		}
	}
	return filterArr
}
