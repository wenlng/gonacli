package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnDoubleTypeCode() string {
	code := ""
	if tools.CheckOS64Unit() {
		code = tools.FormatCodeIndentLn(`const double wg_res_ = (double)wg_data;`, 2)
	} else {
		code = tools.FormatCodeIndentLn(`const double wg_res_ = (double)wg_data;`, 2)
	}

	return code
}

func GenAsyncCallReturnDoubleTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const double wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`
	return code
}

func GenAsyncCallbackArgDoubleTypeCode() string {
	return `Number wg_double_ = Number::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_double_ };`
}
