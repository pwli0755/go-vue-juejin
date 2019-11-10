package conf

import (
	"backend/cache"
	"backend/model"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

//Config   系统配置配置
type Config struct {
	Server `yaml:"server"`
	Gorm   `yaml:"gorm"`
}

type Gorm struct {
	DbLogMode         bool `yaml:"DbLogMode"`         //数据库日志模式，开启true, 关闭false
	DbMaxIdleConns    int  `yaml:"DbMaxIdleConns"`    //最大空闲连接数
	DbMaxOpenConns    int  `yaml:"DbMaxOpenConns"`    //最大连接数
	DbConnMaxLifetime int  `yaml:"DbConnMaxLifetime"` //mysql超时时间
}

type Server struct {
	Port         string `yaml:"port"`
	Debug        bool   `yaml:"debug"`
	MultiDevices bool   `yaml:"multiDevices"`
}

var Conf Config

// 项目相关配置，不含敏感信息
func GetConfig(conf *Config) {
	file, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		panic(err)
	}
}

// init方法会自动执行
func init() {
	// 读取项目配置
	GetConfig(&Conf)
	// 读取并加载环境变量, env一般配置敏感信息
	godotenv.Load("./.env")
	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}

	// 连接数据库
	model.GetDbConn(os.Getenv("MYSQL_DSN"),
		Conf.Gorm.DbMaxIdleConns, Conf.Gorm.DbMaxOpenConns, Conf.Gorm.DbConnMaxLifetime)
	cache.Redis()
}
