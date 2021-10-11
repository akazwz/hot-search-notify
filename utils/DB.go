package utils

import (
	"fmt"
	"github.com/akazwz/hot-search-notify/inital"
	_ "github.com/go-sql-driver/mysql"
)

type Contents struct {
	Id       int    `json:"id"`
	Contents string `json:"contents"`
}

func GetAllSubWords() {
	var result []Contents
	inital.GDB.Raw("SELECT * FROM contents").Scan(&result)
	fmt.Println(result)
}
