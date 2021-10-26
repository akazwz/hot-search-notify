package inital

type Conf struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	DBName    string `yaml:"DBName"`
	Host      string `yaml:"host"`
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}
