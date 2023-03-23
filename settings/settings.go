package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*MySqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      string `mapstructure:"port"`
	JwtExpire int64  `mapstructure:"jwt_expire"`
}

type MySqlConfig struct {
	Host         string `mapstructure:"host"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml") // 指定配置文件路径
	//viper.SetConfigName("config")             // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")               // 如果配置文件的名称中没有扩展名，则需要配置此项
	//viper.AddConfigPath("./conf")             // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	//把读取到的配置信息反序列到Conf变量中去
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err: %v\n", err)
	}

	//实时监控配置文件的变化
	viper.WatchConfig()
	//当配置变化后调用的回调函数
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("Config file changed!!")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
