package resync

import (
	"node_addon_go/config"
	args2 "node_addon_go/content/args/argsync"
	"node_addon_go/tools"
	"strings"
)

// 生成处理体
func GenHandleReturnStringCode(method string, args []string, preCode string) string {
	code := tools.FormatCodeIndentLn(`string wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)

	code += preCode
	code += tools.FormatCodeIndentLn(`return String::New(wg_env, wg_res_);`, 2)
	return code
}

// 生成-返回数字型
func GenReturnStringTypeCode(export config.Export) string {
	methodName := export.Name
	args := export.Args

	code := `
// ---------- GenCode ---------- 
Value _` + methodName + `(const CallbackInfo& wg_info) {`
	code += tools.FormatCodeIndentLn(`Env wg_env = wg_info.Env();`, 2)

	c, argNames, preCode := args2.GenArgCode(args)
	code += c

	code += GenHandleReturnStringCode(methodName, argNames, preCode)

	code += tools.FormatCodeIndentLn(`}`, 0)
	return code
}
