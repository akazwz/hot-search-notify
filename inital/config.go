package inital

type Conf struct {
	Username  string  `yaml:"username"`
	Password  string  `yaml:"password"`
	DBName    string  `yaml:"dbName"`
	Host      string  `yaml:"host"`
	AppId     string  `yaml:"appId"`
	AppSecret string  `yaml:"appSecret"`
	Tencent   Tencent `yaml:"tencent"`
}

type Tencent struct {
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
}
