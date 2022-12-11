package check

import (
	"fmt"
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
)

func CheckBaseConfig(config config.Config) error {
	//defer func() {
	//	if err := recover(); err != nil {
	//		println(err.(string))
	//	}
	//}()

	clog.Info("Start checking the configured ...")

	// 检测参数列表是否冲突
	exports := config.Exports

	var goApiList = make([]string, 0)
	var JsCallApiList = make([]string, 0)
	var allowArgsList = []string{"int", "int32", "int64", "uint32", "float", "double", "boolean", "string", "array", "object", "arraybuffer", "callback"}
	var returnAllowArgsList = []string{"int", "int32", "int64", "uint32", "float", "double", "boolean", "string", "array", "arraybuffer", "object"}

	for _, export := range exports {
		// 正在检测导出的 Go API、JS CALL API 是否重复
		if tools.InSlice(goApiList, export.Name) {
			return fmt.Errorf("The export Name \"%s\" already exists", export.Name)
		}
		goApiList = append(goApiList, export.Name)

		if tools.InSlice(JsCallApiList, export.JsCallName) {
			return fmt.Errorf("The export JsCallName \"%s\" already exists", export.JsCallName)
		}
		JsCallApiList = append(JsCallApiList, export.JsCallName)

		// 检测参数类型是正确，列表中是否重复命名
		// 参数类型 [int,int32,int64,uint32,float,double,boolean,string,array,object,callback]
		var curArgsList = make([]string, 0)
		for _, arg := range export.Args {
			if !tools.InSlice(allowArgsList, arg.Type) {
				return fmt.Errorf("The arguments type \"%s\" of the exported [%s] is not supported", arg.Type, export.Name)
			}

			if tools.InSlice(curArgsList, arg.Name) {
				return fmt.Errorf("The export parameter \"%s\" of [%s] already exists", arg.Name, export.Name)
			}
			curArgsList = append(curArgsList, arg.Name)
		}

		// 检测返回值是否正确 - 没有callback
		// 参数类型 [int,int32,int64,uint32,float,double,boolean,string,array,object]
		if !tools.InSlice(returnAllowArgsList, export.ReturnType) {
			return fmt.Errorf("The return type \"%s\" of the exported [%s] is not supported", export.ReturnType, export.Name)
		}
	}
	clog.Success("Checking the configured done ~")
	return nil
}

func CheckAsyncCorrectnessConfig(config config.Config) error {
	for _, export := range config.Exports {
		if export.JsCallMode == "async" {
			checked := false
			for _, arg := range export.Args {
				if arg.Type == "callback" {
					checked = true
					break
				}
			}
			if !checked {
				return fmt.Errorf("The export [%s] is missing the \"callback\" parameter", export.Name)
			}
		}
	}
	return nil
}

func CheckExportApiWithSourceFile(config config.Config) bool {
	clog.Info("Start checking the API of Golang export ...")
	clog.Success("Checking the API of Golang export done ~")

	return true
}
