package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnObjectTypeCode() string {
	return tools.FormatCodeIndentLn(`const string wg_res_ = (char*)wg_data;`, 2)
}

func GenAsyncCallReturnObjectTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const char* wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`
	return code
}

func GenAsyncCallbackArgObjectTypeCode() string {
	code := `Object wg_obj = wg_string_to_object(wg_res_, wg_env);
    napi_value wg_argv[] = { wg_obj };`
	return code
}
