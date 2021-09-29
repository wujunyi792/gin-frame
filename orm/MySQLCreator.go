package orm

import (
	"fmt"
	"github.com/wujunyi792/StuBigData/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type MySQLCreator struct {
}

func (m MySQLCreator) Create(dbConfig DBConfig) (*gorm.DB, error) {
	var userName = dbConfig.AuthConfig.UserName
	var password = dbConfig.AuthConfig.Password
	var dbName = dbConfig.DBName
	var dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName, password, dbConfig.Address, dbName)
	logger.Info.Printf("conn str: %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//ensureMySqlDatabase(orm, dbConfig.DBName) //重连会不能选中数据库
	return db, err
}

func ensureMySqlDatabase(db *gorm.DB, dbName string) {
	err := db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
	if (err != nil) && (!strings.Contains(err.Error(), "1007")) { //有错误且不是数据库已创建
		logger.Error.Fatalln(err)
		return
	}
	err = db.Exec(fmt.Sprintf("USE %s", dbName)).Error
	if err != nil {
		logger.Error.Fatalln(err)
		return
	}
}
