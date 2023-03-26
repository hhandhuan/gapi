package mysql

import (
	"gapi/pkg/conf"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormDefaultLogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	DB = getConn(conf.GetConfig().Mysql)
}

func getConn(c *conf.Mysql) *gorm.DB {
	gdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Dns, // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger:                 gormDefaultLogger.Default.LogMode(gormDefaultLogger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalf("conn mysql client error: %v", err)
		return nil
	}

	db, err := gdb.DB()
	if err != nil {
		log.Fatalf("conn mysql error: %v", err)
		return nil
	}

	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	db.SetConnMaxIdleTime(time.Second * time.Duration(c.ConnMaxIdleTime))
	db.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))

	return gdb
}
