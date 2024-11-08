package main

import (
  "flag"
  "github.com/lucky-finger/desktop/localserver/db"
  "github.com/lucky-finger/desktop/localserver/logger"
)

var (
  // isDev 是否是开发环境
  isDev = false
  // dbFile 数据库文件
  dbFile = "debug.db"
  // logLevel 日志级别
  logLevel = "error"
  // logPath 日志路径
  logPath = "logs"
)

func main() {
  flag.BoolVar(&isDev, "dev", false, "是否是开发环境")
  flag.StringVar(&dbFile, "db", "debug.db", "数据库文件")
  flag.StringVar(&logLevel, "logLevel", "error", "日志级别")
  flag.StringVar(&logPath, "logPath", "logs", "日志路径")
  flag.Parse()

  if err := db.Load(isDev, dbFile); err != nil {

  }

  if err := logger.Load(logLevel, logPath); err != nil {

  }

}
