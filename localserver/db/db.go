package db

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

// _db 数据库连接
var _db *gorm.DB

type writer struct {
}

func (w *writer) Printf(s string, i ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, fmt.Sprintf(s, i...))
}

// Load 初始化数据库
func Load(isDev bool, dbFile string) (err error) {
	if _db, err = gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?_auto_vacuum=1&_synchronous=0", dbFile)), &gorm.Config{
		Logger: logger.New(&writer{}, logger.Config{
			Colorful: true,
			LogLevel: logger.Info,
		}),
	}); err != nil {
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
