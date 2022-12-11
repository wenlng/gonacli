package async

import (
	"fmt"
	"node_addon_go/config"
	"node_addon_go/content/returns/reasync"
	"node_addon_go/tools"
	"strings"
)

// 生成注册代码
func genRegisterCode(name string, jsCallName string, workName string) (string, string) {
	funName := "register_" + strings.ToLower(name)
	code := `
// ---------- GenRegisterAsyncCode ---------- 
void ` + funName + `(Env env, Object exports){
  napi_property_descriptor desc = {"` + jsCallName + `",NULL,` + workName + `,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}`
	return code, funName + "(env, exports);"
}

// 参数解析
func genResultParseCode(returnType string) (string, string) {
	code := ""
	preCode := ""
	if returnType == "string" {
		ccode, pcode := reasync.GenAsyncReturnStringTypeCode()
		code += ccode
		preCode += pcode
	} else if returnType == "boolean" {
		code = reasync.GenAsyncReturnBooleanTypeCode()
	} else if returnType == "int" {
		code += reasync.GenAsyncReturnIntTypeCode("int32")
	} else if returnType == "int32" {
		code += reasync.GenAsyncReturnIntTypeCode("int32")
	} else if returnType == "int64" {
		code += reasync.GenAsyncReturnIntTypeCode("int64")
	} else if returnType == "uint32" {
		code += reasync.GenAsyncReturnIntTypeCode("unit32")
	} else if returnType == "float" {
		code += reasync.GenAsyncReturnFloatTypeCode()
	} else if returnType == "double" {
		code += reasync.GenAsyncReturnDoubleTypeCode()
	} else if returnType == "array" {
		code += reasync.GenAsyncReturnArrayTypeCode()
	} else if returnType == "object" {
		code += reasync.GenAsyncReturnObjectTypeCode()
	} else if returnType == "arraybuffer" {
		code += reasync.GenAsyncReturnArrayBufferTypeCode()
	} else {
		code = tools.FormatCodeIndentLn(`(void)data;`, 2)
	}

	return code, preCode
}

func genStructCallbackCode(export config.Export, structName string) string {
	argLen := len(export.Args)
	code := `
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[` + fmt.Sprintf("%d", argLen) + `];
} ` + structName + `;`
	return code
}

// 生成-返回数字型
func GenAsyncCallbackCode(export config.Export) (string, string) {
	methodName := export.Name
	//args := export.Args
	workName := "wg_work_" + strings.ToLower(methodName)
	workCompleteName := "wg_work_complete_" + strings.ToLower(methodName)
	executeworkName := "wg_execute_work" + strings.ToLower(methodName)
	jsCallbackName := "wg_js_callback_" + strings.ToLower(methodName)
	structDataName := "WgAddonData" + methodName

	code := `
// [` + methodName + `] +++++++++++++++++++++++++++++++++ start`
	code += genStructCallbackCode(export, structDataName)
	code += genJsCallbackCode(export, jsCallbackName)
	code += genExecuteWorkCode(export, executeworkName, structDataName)
	code += genWorkCompleteCode(workCompleteName, structDataName)
	code += genWorkThreadCode(
		export,
		workName,
		workCompleteName,
		executeworkName,
		jsCallbackName,
		structDataName,
	)

	// 生成注册代码
	rCode, rFunc := genRegisterCode(methodName, export.JsCallName, workName)
	code += rCode

	code += `
// [` + methodName + `]+++++++++++++++++++++++++++++++++ end`

	registerCode := tools.FormatCodeIndentLn(rFunc, 2)
	return code, registerCode
}
