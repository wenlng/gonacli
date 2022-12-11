package reasync

import (
	"node_addon_go/tools"
	"strings"
)

func GenAsyncReturnIntTypeCode(varType string) string {
	code := ""

	// 	int32 int64 uint32
	if varType == "int64" {
		code = tools.FormatCodeIndentLn(`const long long int wg_res_ = (long long int)(intptr_t)wg_data;`, 2)
	} else if varType == "uint32" {
		code = tools.FormatCodeIndentLn(`const unsigned long wg_res_ = (unsigned long)(intptr_t)wg_data;`, 2)
	} else {
		code = tools.FormatCodeIndentLn(`const int wg_res_ = (int)(intptr_t)wg_data;`, 2)
	}

	return code
}

func GenAsyncCallReturnIntTypeCode(methodName string, argNames []string, varType string) string {
	code := ""
	// 	int64 uint32 int32
	//  int64_t uint32_t int32_t
	if varType == "int64" {
		code = `
  // -------- genHandlerCode
  const const long long wg_res__ = ` + methodName + `(` + strings.Join(argNames, ",") + `);
  void * wg_res_ = (void *)(intptr_t)wg_res__;`
	} else if varType == "uint32" {
		code = `
  // -------- genHandlerCode
  const const unsigned long wg_res__ = ` + methodName + `(` + strings.Join(argNames, ",") + `);
  void * wg_res_ = (void *)(intptr_t)wg_res__;`
	} else {
		code = `
  // -------- genHandlerCode
  const int wg_res__ = ` + methodName + `(` + strings.Join(argNames, ",") + `);
  void * wg_res_ = (void *)(intptr_t)wg_res__;`
	}

	return code
}

func GenAsyncCallbackArgIntTypeCode() string {
	return `Number wg_int_ = Number::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_int_ };`
}
