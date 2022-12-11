package async

import (
	"fmt"
	"node_addon_go/config"
	"node_addon_go/content/args/argasync"
	"node_addon_go/content/returns/reasync"
	"node_addon_go/tools"
)

// 生成参数
func genHandlerArgCode(arg config.Arg, index int) (string, string) {
	code := ""
	preCode := ""
	sIndex := fmt.Sprintf("%d", index)

	if arg.Type == "string" {
		cCode, pCode := argasync.GenAsyncStringArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "boolean" {
		code += argasync.GenAsyncBooleanArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "int" {
		code += argasync.GenAsyncIntArgTypeCode(arg.Name, sIndex, "int32")
	} else if arg.Type == "int32" {
		code += argasync.GenAsyncIntArgTypeCode(arg.Name, sIndex, "int32")
	} else if arg.Type == "int64" {
		code += argasync.GenAsyncIntArgTypeCode(arg.Name, sIndex, "int64")
	} else if arg.Type == "uint32" {
		code += argasync.GenAsyncIntArgTypeCode(arg.Name, sIndex, "uint32")
	} else if arg.Type == "float" {
		code += argasync.GenAsyncFloatArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "double" {
		code += argasync.GenAsyncDoubleArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "array" {
		cCode, pCode := argasync.GenAsyncArrayArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "object" {
		cCode, pCode := argasync.GenAsyncObjectArgTypeCode(arg.Name, sIndex, 2)
		code += cCode
		preCode += pCode
	} else if arg.Type == "arraybuffer" {
		cCode, pCode := argasync.GenAsyncArrayBufferArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "callback" {
		cCode, pCode := argasync.GenAsyncFuncitonArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	}

	return code, preCode
}

// 业务逻辑
func genHandlerCode(export config.Export) (string, string) {
	methodName := export.Name
	returnType := export.ReturnType

	code := ""
	preCode := ""
	var argNames = make([]string, 0)
	for index, arg := range export.Args {
		cCode, pCode := genHandlerArgCode(arg, index)
		code += cCode
		preCode += pCode
		argNames = append(argNames, arg.Name)
	}

	if returnType == "string" {
		code += reasync.GenAsyncCallReturnStringTypeCode(methodName, argNames)
	} else if returnType == "boolean" {
		code += reasync.GenAsyncCallReturnBooleanTypeCode(methodName, argNames)
	} else if returnType == "int" {
		code += reasync.GenAsyncCallReturnIntTypeCode(methodName, argNames, "int32")
	} else if returnType == "int32" {
		code += reasync.GenAsyncCallReturnIntTypeCode(methodName, argNames, "int32")
	} else if returnType == "int64" {
		code += reasync.GenAsyncCallReturnIntTypeCode(methodName, argNames, "int64")
	} else if returnType == "uint32" {
		code += reasync.GenAsyncCallReturnIntTypeCode(methodName, argNames, "uint32")
	} else if returnType == "float" {
		code += reasync.GenAsyncCallReturnFloatTypeCode(methodName, argNames)
	} else if returnType == "double" {
		code += reasync.GenAsyncCallReturnDoubleTypeCode(methodName, argNames)
	} else if returnType == "array" {
		code += reasync.GenAsyncCallReturnArrayTypeCode(methodName, argNames)
	} else if returnType == "object" {
		code += reasync.GenAsyncCallReturnObjectTypeCode(methodName, argNames)
	} else if returnType == "arraybuffer" {
		code += reasync.GenAsyncCallReturnArrayBufferTypeCode(methodName, argNames)
	} else {
		code += tools.FormatCodeIndentLn(`const void* res =`+methodName+`();`, 4)
	}

	return code, preCode
}
