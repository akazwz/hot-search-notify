package utils

import (
	"github.com/akazwz/hot-search-notify/inital"
	"log"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

func SendMsg(userUUID, openId string) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     inital.CFG.AppId,
		AppSecret: inital.CFG.AppSecret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	sub := mini.GetSubscribe()
	data := make(map[string]*subscribe.DataItem)
	data["phrase1"] = &subscribe.DataItem{
		Value: "test",
		Color: "",
	}
	data["phrase1"] = &subscribe.DataItem{
		Value: "吃饭了吗",
		Color: "",
	}
	data["date2"] = &subscribe.DataItem{
		Value: "2019-12-11 11:00:00",
		Color: "",
	}
	data["phrase3"] = &subscribe.DataItem{
		Value: "点击查看",
		Color: "",
	}
	msg := &subscribe.Message{
		ToUser:     openId,
		TemplateID: "XV16ZyG6Af_gG8D4qg7M17Fw23m_zYWNo689XpJKYQE",
		Data:       data,
	}
	err := sub.Send(msg)
	if err != nil {
		log.Println("send error")
		log.Println(err)
		return
	}
	//inital.GDB.Create()

}
