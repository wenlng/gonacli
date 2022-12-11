package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnArrayTypeCode() string {
	return tools.FormatCodeIndentLn(`const string wg_res_ = (char*)wg_data;`, 2)
}

func GenAsyncCallReturnArrayTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const char* wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`
	return code
}

func GenAsyncCallbackArgArrayTypeCode() string {
	code := `Array wg_arr = wg_string_to_array(wg_res_, wg_env);
    napi_value wg_argv[] = { wg_arr };`
	return code
}
