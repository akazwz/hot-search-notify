package inital

import "github.com/spf13/viper"

func InitViper() (config *viper.Viper) {
	config = viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	// 读取配置
	if err := config.ReadInConfig(); err != nil {
		panic(err)
		return nil
	}

	// 解析配置
	if err := config.Unmarshal(&CFG); err != nil {
		panic(err)
		return nil
	}
	return
}
