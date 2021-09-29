package example

import (
	"github.com/wujunyi792/ginFrame/logger"
	"github.com/wujunyi792/ginFrame/model/dbModel"
	"github.com/wujunyi792/ginFrame/orm"
	"gorm.io/gorm"
	"sync"
)

type ExampleManage struct {
	mDB     *orm.MainGORM
	sDBLock sync.RWMutex
}

var loginLogManage *ExampleManage = nil

func (manager *ExampleManage) getGOrmDB() *gorm.DB {
	return manager.mDB.GetDB()
}

func (manager *ExampleManage) atomicDBOperation(op func()) {
	manager.sDBLock.Lock()
	op()
	manager.sDBLock.Unlock()
}

func GetExampleManage() *ExampleManage {
	if loginLogManage == nil {
		//创建连接
		var db = orm.MustCreateStuGOrm("User")
		err := db.GetDB().AutoMigrate(&dbModel.Example{})
		if err != nil {
			logger.Error.Fatalln(err)
			return nil
		}
		loginLogManage = &ExampleManage{mDB: db}
	}
	return loginLogManage
}
