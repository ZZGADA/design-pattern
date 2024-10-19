package client

import (
	"fmt"
	_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"sync"
	"syncache/conf"
	"time"
)

var (
	MysqlInstance *Mysql
)

func init() {
	MysqlInstance = &Mysql{}
}

type Mysql struct {
	sync.Once
	client *gorm.DB
}

type dataSourceMysql interface {
	Get(config conf.Config) *gorm.DB
}

// Get 获取客户端
func (mysql *Mysql) Get(config conf.Config) *gorm.DB {
	// 追加单例锁
	mysql.Do(func() {
		const mysqlConnectStr string = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		mysqlConf := config.MysqlConfig
		dsn := fmt.Sprintf(mysqlConnectStr,
			mysqlConf.Username,
			mysqlConf.Password,
			mysqlConf.Host,
			mysqlConf.Port,
			mysqlConf.DataBase)

		client, err := gorm.Open(_mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "",   // 表前缀
				SingularTable: true, // 禁用表名复数
			}})

		if err != nil {
			panic(err)
		}

		sqlDB, _ := client.DB()
		// SetMaxIdleConnections 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConnections 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(10)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(10 * time.Second)

		mysql.client = client
	})

	return mysql.client
}
