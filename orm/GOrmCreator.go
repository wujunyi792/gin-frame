package orm

import (
	"github.com/wujunyi792/ginFrame/logger"
	"gorm.io/gorm"
)

type Creator interface {
	Create(dbConfig DBConfig) (*gorm.DB, error)
}

var sSelectorMap = make(map[string]Creator)

func GetCreatorByType(dbType string) Creator {
	return sSelectorMap[dbType]
}

func init() {
	logger.Info.Println("init orm creators")
	sSelectorMap["mysql"] = MySQLCreator{}
}
