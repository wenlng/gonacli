package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnFloatTypeCode() string {
	code := tools.FormatCodeIndentLn(`const float wg_res_ = (float)wg_data;`, 2)

	return code
}

func GenAsyncCallReturnFloatTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const float wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`

	return code
}

func GenAsyncCallbackArgFloatTypeCode() string {
	return `Number wg_float_ = Number::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_float_ };`
}
