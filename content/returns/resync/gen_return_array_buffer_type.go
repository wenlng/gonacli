package resync

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/args/argsync"
	"github.com/wenlng/gonacli/tools"
	"strings"
)

// 生成处理体
func GenHandleReturnArrayBufferCode(method string, args []string, preCode string) string {
	// 转换成数组buffer
	code := `
  void * wg_res_ = ` + method + `(` + strings.Join(args, ",") + `);
  byte *wg_ab = (byte*) wg_res_;
  size_t wg_ab_length = strlen((char *)wg_ab);
  ArrayBuffer wg_arr_buffer = ArrayBuffer::New(wg_env, wg_ab, wg_ab_length);`

	code += preCode
	code += tools.FormatCodeIndentLn(`return wg_arr_buffer;`, 2)
	return code
}

// 生成-返回数字型
func GenReturnArrayBufferTypeCode(export config.Export) string {
	methodName := export.Name
	args := export.Args

	code := `
// ---------- GenCode ---------- 
Value _` + methodName + `(const CallbackInfo& wg_info) {`
	code += tools.FormatCodeIndentLn(`Env wg_env = wg_info.Env();`, 2)

	c, argNames, endCode := argsync.GenArgCode(args)
	code += c

	code += GenHandleReturnArrayBufferCode(methodName, argNames, endCode)

	code += tools.FormatCodeIndentLn(`}`, 0)
	return code
}
