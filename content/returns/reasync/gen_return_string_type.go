package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnStringTypeCode() (string, string) {
	code := tools.FormatCodeIndentLn(`const string wg__res_ = (char*)wg_data;`, 2)

	pCode := tools.FormatCodeIndentLn(`napi_value wg_res_ = String::New(wg_env, wg__res_);`, 4)
	//endCode := tools.FormatCodeIndentLn(`delete [] _res_`, 0)
	return code, pCode
}

func GenAsyncCallReturnStringTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const void* wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`
	return code
}

func GenAsyncCallbackArgStringTypeCode() string {
	//return `napi_value wg_string_ = String::New(wg_env, wg_res_);
	//napi_value wg_argv[] = { wg_string_ };`
	return `napi_value wg_argv[] = { wg_res_ };`
}
