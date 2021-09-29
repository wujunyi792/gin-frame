package config

import (
	"github.com/pelletier/go-toml"
	"github.com/wujunyi792/ginFrame/logger"
	"reflect"
)

var sInstanceMap = make(map[string]*TConfig)

type TConfig struct {
	mRootTree *toml.Tree
}

func GetConfig(alias string) *TConfig {
	var config = sInstanceMap[alias]
	if config == nil {
		logger.Error.Fatalln("must call InitConfig first")
		return nil
	}
	return config
}

func InitConfig(alias string, configPath string, fatalWhenError bool) {
	rootTree, err := toml.LoadFile(configPath)
	if err != nil {
		if fatalWhenError {
			logger.Error.Fatalln(err)
		} else {
			logger.Error.Println(err)
		}
		return
	}
	sInstanceMap[alias] = &TConfig{mRootTree: rootTree}
}

func (config TConfig) GetRootTree() *toml.Tree {
	return config.mRootTree
}

// 在包含模块名的父节点上获取功能节点的struct或值，$outRef必须传引用
func (config TConfig) Unmarshall(moduleName string, nodeName string, outRef interface{}) error {
	var moduleTree = config.GetRootTree().Get(moduleName).(*toml.Tree)
	var treeType = reflect.TypeOf(moduleTree)
	node := moduleTree.Get(nodeName)
	if reflect.TypeOf(node).ConvertibleTo(treeType) {
		var nodeTree = node.(*toml.Tree)
		err := nodeTree.Unmarshal(outRef)
		return err
	} else {
		reflect.ValueOf(outRef).Elem().Set(reflect.ValueOf(node))
		return nil
	}

}
