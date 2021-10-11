package utils

import (
	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func GetAllSubWords() []string {
	var allWords string
	inital.GDB.Raw("SELECT GROUP_CONCAT(sub_word) FROM all_sub_words").Scan(&allWords)
	return strings.Split(allWords, ",")
}

func GetFilterUserInfo() {

}

func CreateSubWords() {
	err := inital.GDB.Create(&model.Sub{
		UserId:   111,
		Phone:    "15153953308",
		SubWords: []byte(`["赵文卓", "赵文卓", "张杰"]`),
	}).Error
	if err != nil {
		log.Println(err)
	}
}

func InsertAllSubWords() {
	words := []string{"赵文卓", "张杰"}
	for i := 0; i < len(words); i++ {
		err := inital.GDB.Create(&model.AllSubWords{
			SubWord: words[i],
		}).Error
		if err != nil {
			log.Println(err)
		}
	}
}
