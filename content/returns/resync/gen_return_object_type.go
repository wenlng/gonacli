package resync

import (
	"node_addon_go/config"
	args2 "node_addon_go/content/args/argsync"
	"node_addon_go/tools"
	"strings"
)

// 生成处理体
func GenHandleReturnObjectCode(method string, args []string, preCode string) string {
	code := tools.FormatCodeIndentLn(`string wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)

	code += tools.FormatCodeIndentLn(`Object wg_obj_ = wg_string_to_object(wg_res_, wg_env);`, 2)
	code += preCode
	code += tools.FormatCodeIndentLn(`return wg_obj_;`, 2)
	return code
}

// 生成-返回数字型
func GenReturnObjectTypeCode(export config.Export) string {
	methodName := export.Name
	args := export.Args

	code := `
// ---------- GenCode ---------- 
Value _` + methodName + `(const CallbackInfo& wg_info) {`
	code += tools.FormatCodeIndentLn(`Env wg_env = wg_info.Env();`, 2)

	c, argNames, endCode := args2.GenArgCode(args)
	code += c

	code += GenHandleReturnObjectCode(methodName, argNames, endCode)

	code += tools.FormatCodeIndentLn(`}`, 0)
	return code
}
