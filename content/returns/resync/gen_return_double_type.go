package resync

import (
	"github.com/wenlng/gonacli/config"
	args2 "github.com/wenlng/gonacli/content/args/argsync"
	"github.com/wenlng/gonacli/tools"
	"strings"
)

// 生成处理体
func GenHandleReturnDoubleCode(method string, args []string) string {
	code := tools.FormatCodeIndentLn(`double wg_res_ = `+method+`(`+strings.Join(args, ",")+`);`, 2)
	code += tools.FormatCodeIndentLn(`return Number::New(wg_env, wg_res_);`, 2)
	return code
}

// 生成-返回数字型
func GenReturnDoubleTypeCode(export config.Export) string {
	methodName := export.Name
	args := export.Args

	code := `
// ---------- GenCode ---------- 
Value _` + methodName + `(const CallbackInfo& wg_info) {`
	code += tools.FormatCodeIndentLn(`Env wg_env = wg_info.Env();`, 2)

	c, argNames, _ := args2.GenArgCode(args)
	code += c

	code += GenHandleReturnDoubleCode(methodName, argNames)

	code += tools.FormatCodeIndentLn(`}`, 0)
	return code
}
