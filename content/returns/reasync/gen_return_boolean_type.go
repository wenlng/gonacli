package reasync

import (
	"node_addon_go/tools"
	"strings"
)

func GenAsyncReturnBooleanTypeCode() string {
	return tools.FormatCodeIndentLn(`const bool wg_res_ = (bool)wg_data;`, 2)
}

func GenAsyncCallReturnBooleanTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const bool wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`

	return code
}

func GenAsyncCallbackArgBooleanTypeCode() string {
	return `Boolean wg_bool_ = Boolean::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_bool_ };`
}
