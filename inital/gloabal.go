package inital

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	VP  *viper.Viper
	CFG Conf
	GDB *gorm.DB
)
