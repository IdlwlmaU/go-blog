package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name          string `mapstructure:"name"`
	Mode          string `mapstructure:"mode"`
	Version       string `mapstructure:"version"`
	Port          int    `mapstructure:"port"`
	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*Viewer       `mapstructure:"viewer"`
	*SystemConfig `mapstructure:"system_config"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_size"`
	MaxBackUps int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type Viewer struct {
	Title       string   `mapstructure:"title"`
	Description string   `mapstructure:"description"`
	Logo        string   `mapstructure:"logo"`
	Navigation  []string `mapstructure:"navigation"`
	Bilibili    string   `mapstructure:"bilibili"`
	Zhihu       string   `mapstructure:"zhihu"`
	Avatar      string   `mapstructure:"avatar"`
	UserName    string   `mapstructure:"user_name"`
	UserDesc    string   `mapstructure:"user_desc"`
}
type SystemConfig struct {
	CdnURL          string `mapstructure:"cdn_url"`
	QiniuAccessKey  string `mapstructure:"qiniu_access_key"`
	QiniuSecretKey  string `mapstructure:"qiniu_secret_key"`
	Valine          bool   `mapstructure:"valine"`
	ValineAppid     string `mapstructure:"valine_app_id"`
	ValineAppkey    string `mapstructure:"valine_app_key"`
	ValineServerURL string `mapstructure:"valine_server_url"`
}

// Conf 全局变量，用来保存程序所有的配置信息
var Conf = new(AppConfig)

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper.ReadInConfig() failed, err:", err)
		return
	}
	// 反序列化
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Println("viper.Unmarshal failed, err:", err)
		return
	}
	//fmt.Printf("%#v, %#v", Conf.Viewer, Conf.SystemConfig)
	//fmt.Println(Conf.Viewer)
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		err = viper.Unmarshal(Conf)
		if err != nil {
			fmt.Println("viper.Unmarshal failed, err:", err)
			return
		}
	})
	return
}
