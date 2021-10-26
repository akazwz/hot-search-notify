package utils

import (
	"log"
	"strings"

	"github.com/akazwz/hot-search-notify/inital"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllSubWords() []string {
	var allWords string
	inital.GDB.Raw("SELECT GROUP_CONCAT(sub_words) FROM sub").Scan(&allWords)
	log.Println(allWords)
	allWords = strings.ReplaceAll(allWords, "[", "")
	allWords = strings.ReplaceAll(allWords, "]", "")
	allWords = strings.ReplaceAll(allWords, `"`, "")
	arrBefore := strings.Split(allWords, ",")
	var arrReturn []string
	for i := 0; i < len(arrBefore); i++ {
		arrReturn = append(arrReturn, strings.TrimSpace(arrBefore[i]))
	}
	return arrReturn
}
