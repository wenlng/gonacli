package async

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/returns/reasync"
	"github.com/wenlng/gonacli/tools"
)

// js 回调参数
func genJsCallbackArgs(export config.Export) string {
	returnType := export.ReturnType

	code := ""
	if returnType == "string" {
		code += reasync.GenAsyncCallbackArgStringTypeCode()
	} else if returnType == "boolean" {
		code += reasync.GenAsyncCallbackArgBooleanTypeCode()
	} else if returnType == "int" {
		code += reasync.GenAsyncCallbackArgIntTypeCode()
	} else if returnType == "int32" {
		code += reasync.GenAsyncCallbackArgIntTypeCode()
	} else if returnType == "int64" {
		code += reasync.GenAsyncCallbackArgIntTypeCode()
	} else if returnType == "uint32" {
		code += reasync.GenAsyncCallbackArgIntTypeCode()
	} else if returnType == "float" {
		code += reasync.GenAsyncCallbackArgFloatTypeCode()
	} else if returnType == "double" {
		code += reasync.GenAsyncCallbackArgDoubleTypeCode()
	} else if returnType == "array" {
		code += reasync.GenAsyncCallbackArgArrayTypeCode()
	} else if returnType == "object" {
		code += reasync.GenAsyncCallbackArgObjectTypeCode()
	} else if returnType == "arraybuffer" {
		code += reasync.GenAsyncCallbackArgArrayBufferTypeCode()
	} else {
		code += tools.FormatCodeIndentLn(`napi_value argv[] = {};`, 4)
	}

	return code
}

// js 回调
func genJsCallbackCode(export config.Export, jsCallbackName string) string {
	parseCode, parsePreCode := genResultParseCode(export.ReturnType)
	code := `
// ------------ genJsCallbackCode
static void ` + jsCallbackName + `(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;` + parseCode + `
  if (wg_env != NULL) {` + parsePreCode + `
    ` + genJsCallbackArgs(export) + `
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}`
	return code
}
