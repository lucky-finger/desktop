package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// _db 数据库连接
var _db *gorm.DB

// Load 初始化数据库
func Load(isDev bool, dbFile string) (err error) {
	if _db, err = gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?_auto_vacuum=1&_synchronous=0", dbFile)), &gorm.Config{}); err != nil {
		return err
	}
	if isDev {
		_db = _db.Debug()
	}
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return _db
}
