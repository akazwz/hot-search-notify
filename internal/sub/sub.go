package sub

import (
	"encoding/json"
	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/akazwz/hot-search-notify/internal/utils"
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
	wordsAndContentsMap, subWords := GetFilterSubWordsAndContents(allSubWordsArr)
	log.Println(wordsAndContentsMap)
	log.Println(subWords)
	var allSubModels [][]model.Sub
	for i := 0; i < len(subWords); i++ {
		var subModels []model.Sub
		inital.GDB.Raw(`SELECT * FROM sub WHERE JSON_CONTAINS(sub_words, ?)`, "\""+subWords[i]+"\"").Scan(&subModels)
		allSubModels = append(allSubModels, subModels)
	}
	for i := 0; i < len(allSubModels); i++ {
		subs := allSubModels[i]
		for j := 0; j < len(subs); j++ {
			log.Println(subs[j].SubWords)
		}
	}
}

// GetFilterSubWordsAndContents 传入所有的订阅词汇, 返回符合的订阅词汇和热搜内容
func GetFilterSubWordsAndContents(subWordsArr []string) (map[string][]string, []string) {
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
	log.Println(searches)
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
				// map key 存在
				if _, ok := filterWordContentsMap[filterWordsArr[j]]; ok {
					filterWordContentsMap[filterWordsArr[j]] = append(filterWordContentsMap[filterWordsArr[j]], contentsArr[i])
				} else {
					filterWordContentsMap[filterWordsArr[j]] = []string{contentsArr[i]}
				}
			}
		}
	}
	return filterWordContentsMap, filterWordsArr
}

// 获取热搜包含的订阅词汇
func getFilterWords(subWordsArr, contentsArr []string) []string {
	filterArr := make([]string, 0)
	for i := 0; i < len(contentsArr); i++ {
		for j := 0; j < len(subWordsArr); j++ {
			contains := strings.Contains(contentsArr[i], subWordsArr[j])
			if contains {
				filterArr = append(filterArr, subWordsArr[j])
			}
		}
	}
	return RemoveRepByLoop(filterArr)
}

func RemoveRepByLoop(slc []string) []string {
	var result []string // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
