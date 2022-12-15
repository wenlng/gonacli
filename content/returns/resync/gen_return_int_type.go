package resync

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/args/argsync"
	"github.com/wenlng/gonacli/tools"
	"strings"
)

// 生成处理体
func GenHandleReturnIntCode(method string, args []string, varType string, endCode string) string {
	code := ""

	// 	int32 int64 uint32
	if varType == "int64" {
		code = tools.FormatCodeIndentLn(`long long int wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)
	} else if varType == "uint32" {
		code = tools.FormatCodeIndentLn(`unsigned long wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)
	} else {
		code = tools.FormatCodeIndentLn(`int wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)
	}

	code += endCode
	code += tools.FormatCodeIndentLn(`return Number::New(wg_env, wg_res_);`, 2)
	return code
}

// 生成-返回数字型
func GenReturnIntTypeCode(export config.Export, varType string) string {
	methodName := export.Name
	args := export.Args

	code := `
// ---------- GenCode ---------- 
Value _` + methodName + `(const CallbackInfo& wg_info) {`
	code += tools.FormatCodeIndentLn(`Env wg_env = wg_info.Env();`, 2)

	c, argNames, endCode := argsync.GenArgCode(args)
	code += c

	code += GenHandleReturnIntCode(methodName, argNames, varType, endCode)

	code += tools.FormatCodeIndentLn(`}`, 0)
	return code
}
