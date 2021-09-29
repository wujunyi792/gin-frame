package config

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"github.com/wujunyi792/StuBigData/logger"
	"regexp"
	"strconv"
)

var Config map[string]string

// 配置文件操作句柄
var gConfHandler *goconfig.ConfigFile

func init() {
	//加载配置，注册变量
	//在开发时，只需要更改配置文件名即可
	conf, err := goconfig.LoadConfigFile("config/config.ini")
	gConfHandler = conf
	if err != nil {
		logger.Error.Fatalln(err)
	}

	Config = make(map[string]string)

	for _, configName := range conf.GetSectionList() {
		sec, err := conf.GetSection(configName)
		if err != nil {
			logger.Error.Fatalln(err)
		}

		for key, value := range sec {
			name := fmt.Sprintf("%s_%s", configName, key)
			Config[name] = value
		}
	}

	logger.Info.Printf("config=%v", Config)
}

// @description
// 按段名和键名读取配置值，并自动解析为相应的数据类型存入一个接口类型的值
// "true" | "false" => bool
// ^[1-9]*$ => uint64
// string => string
// @description
// @param	section		string	"段，即config.ini中[]内的字段"
// @param	key			string	"段中的配置项键名"
// @return	res 		interface{}	"读取结果"
func ReadConfValue(section string, key string) interface{} {
	var res interface{}
	value, err := gConfHandler.GetValue(section, key)
	if err != nil {
		logger.Error.Fatalln(err)
	} else {
		if value == "true" || value == "false" {
			res, _ = strconv.ParseBool(value)
		} else if match, _ := regexp.MatchString("^[0-9]*$", value); match {
			res, _ = strconv.ParseUint(value, 10, 64)
		} else {
			res = value
		}

	}
	return res
}
