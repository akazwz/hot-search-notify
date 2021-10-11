package sub

import (
	"encoding/json"
	"github.com/akazwz/hot-search-notify/internal/utils"
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

func NotifySub() {
	allSubWordsArr := utils.GetAllSubWords()
	log.Println("所有订阅词汇:", allSubWordsArr)
	wordsAndContents := GetFilterSubWordsAndContents(allSubWordsArr)
	log.Println(wordsAndContents)
}

// GetFilterSubWordsAndContents 传入所有的订阅词汇, 返回符合的订阅词汇和热搜内容
func GetFilterSubWordsAndContents(subWordsArr []string) map[string][]string {
	// 请求接口,获取当前热搜
	response, err := http.Get("https://hs.hellozwz.com/hot-searches/current")
	if err != nil {
		log.Println("请求失败")
	}
	if response.StatusCode != http.StatusOK {
		log.Println("请求失败")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read error")
	}
	// json 结构到 res
	res := Res{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Println("json error")
	}
	// 热搜数组
	searches := res.Data.Searches
	// 储存热搜内容数组
	contentsArr := make([]string, 0)
	for i := 0; i < len(searches); i++ {
		content := searches[i].Content
		contentsArr = append(contentsArr, content)
	}
	// 热搜包含的订阅词汇
	filterWordsArr := getFilterWords(subWordsArr, contentsArr)
	log.Println("符合条件订阅词汇:", filterWordsArr)
	// 符合订阅的热搜内容
	filterWordContentsMap := make(map[string][]string)
	for i := 0; i < len(contentsArr); i++ {
		for j := 0; j < len(filterWordsArr); j++ {
			isContains := strings.Contains(contentsArr[i], filterWordsArr[j])
			if isContains {
				filterWordContentsMap[filterWordsArr[j]] = []string{contentsArr[i]}
			}
		}
	}

	return filterWordContentsMap
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
