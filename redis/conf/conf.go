package conf

import (
	"code.be.staff.com/staff/StaffGo/public/zap"
	"github.com/go-ini/ini"
	redis "gopkg.in/redis.v5"
)

var (
	// Conf global config object
	Conf *Config
)

// Config 配置
type Config struct {
	Base struct {
		LogConfig string // 日志配置文件路径
	}
	Redis redis.Options
	Zap   zap.Conf
}

func init() {
	Conf = &Config{}
	// 默认空路径
	Conf.Base.LogConfig = ""
}

// Init 初始化 加载并解析配置文件到 Conf 对象
func Init(confPath string) error {
	if err := ini.MapTo(Conf, confPath); err != nil {
		return err
	}
	return nil
}

// Reload 重新加载 重载配置文件到 Conf 对象
func Reload(confPath string) error {
	tmp := &Config{}
	if err := ini.MapTo(tmp, confPath); err != nil {
		return err
	}
	Conf = tmp
	return nil
}
