/**
 * @Author: Cc
 * @Description: mysql client
 * @File: mysql
 * @Version: 1.0.0
 * @Date: 2023/7/21 14:48
 * @Software : GoLand
 */

package mysql

import (
	"log"
	"os"
	"time"

	"github.com/Cc360428/HelpPackage/mysqlc"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var client *gorm.DB

func NewClient() {
	var (
		address  = "127.0.0.1"
		port     = "3306"
		username = "root"
		password = "test"
		database = "telegram_robot"
	)

	dbConfig := new(gorm.Config)

	dbConfig.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "t_",
		SingularTable: true,
		NameReplacer:  nil,
		NoLowerCase:   false,
	}

	dbConfig.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Warn, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		})
	start, err := mysqlc.InitStart(address, port, username, password, database, dbConfig)
	if err != nil {
		return
	}
	err = start.AutoMigrate(&Platform{}, &Game{}, &Keyboard{})
	if err != nil {
		log.Println("Auto :", err.Error())
	}
	client = start
}
