package utils

import (
	"log"
	"time"

	"github.com/akazwz/hot-search-notify/inital"
	"github.com/akazwz/hot-search-notify/model"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func SendMiniMsg(userUUID, openId string) {
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
		Page:       "pages/sub/index",
	}
	err := sub.Send(msg)
	if err != nil {
		log.Println("通知失败")
		log.Println(err)
		return
	}
	var notify model.Notify
	// 数据库更新通知次数和上次通知时间
	err = inital.GDB.Where("user_uuid = ?", userUUID).First(&notify).Updates(&model.Notify{
		NotifyCount:    notify.NotifyCount + 1,
		AllNotifyCount: notify.AllNotifyCount + 1,
		LastNotify:     time.Now(),
	}).Error
	if err != nil {
		log.Println("通知成功,存入数据库失败")
		return
	}
}

func SendSMS(phone string) {
	credential := common.NewCredential(inital.CFG.Tencent.SecretId, inital.CFG.Tencent.SecretKey)

	cpf := profile.NewClientProfile()
	cpf.SignMethod = "HmacSHA1"

	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	request := sms.NewSendSmsRequest()

	request.SmsSdkAppId = common.StringPtr("1400576425")
	request.SignName = common.StringPtr("赵文卓工作学习")
	request.SenderId = common.StringPtr("")
	request.ExtendCode = common.StringPtr("")
	request.TemplateParamSet = common.StringPtrs([]string{})
	request.TemplateId = common.StringPtr("1131592")
	request.PhoneNumberSet = common.StringPtrs([]string{phone})

	_, err := client.SendSms(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		log.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		log.Println(err)
		return
	}
	return
}
