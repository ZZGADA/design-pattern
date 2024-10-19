package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	Dft dfter
)

// init 初始化变量 初始化全局使用的config类
func init() {
	Dft = &dft{}
}

type MysqlConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Port     string `yaml:"port" mapstructure:"port"`
	DataBase string `yaml:"database" mapstructure:"database"`
}

type RedisConfig struct {
	Host           string `yaml:"host" mapstructure:"host"`
	Password       string `yaml:"password" mapstructure:"password"`
	Port           string `yaml:"port" mapstructure:"port"`
	DataBase       int    `yaml:"database" mapstructure:"database"`
	MaxActiveConns int    `yaml:"max_active_conns" mapstructure:"max_active_conns"`
}

type Config struct {
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
}

type dft struct {
	sync.Once
	Config Config
}

// dfter 私有接口 包内可见
type dfter interface {
	Get() Config
}

func (dtf *dft) Get() Config {
	dtf.Do(func() {
		// 单次锁 方式多个go routine 同时初始化
		// 配置读取yaml 文件
		viper.SetConfigName("application") // 配置文件名称(无扩展名)
		viper.SetConfigType("yaml")        // 或viper.SetConfigType("YAML")
		viper.AddConfigPath("./conf")      // 配置文件路径
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		mysqlConfig := MysqlConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			DataBase: viper.GetString("db.database"),
			Username: viper.GetString("db.username"),
			Password: viper.GetString("db.password"),
		}

		redisConfig := RedisConfig{
			Host:           viper.GetString("redis.host"),
			Port:           viper.GetString("redis.port"),
			DataBase:       viper.GetInt("redis.database"),
			Password:       viper.GetString("redis.password"),
			MaxActiveConns: viper.GetInt("redis.max_active_conns"),
		}

		config := Config{MysqlConfig: mysqlConfig, RedisConfig: redisConfig}
		dtf.Config = config
	})

	return dtf.Config
}
