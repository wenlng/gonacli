package argasync

import (
	"fmt"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/validate"
)

// 生成校验代码
func genCheckInputArgCode(code string, requireCode string, validateCode string, index int) string {
	sIndex := fmt.Sprintf("%d", index)
	return `
  Value wg_v` + sIndex + ` = Value(wg_env, wg_args[` + sIndex + `]);` + requireCode + validateCode + code
}

// 生成参数
func GenParseInputArgCode(arg config.Arg, index int) (string, string) {
	code := ""
	preCode := ""
	validateCode := ""
	beforeCode := ""
	sIndex := fmt.Sprintf("%d", index)

	// 生成校验
	if arg.IsRequire {
		beforeCode += validate.GenAsyncRequireWithIndexCode(index)
	}

	if arg.Type == "string" {
		validateCode += validate.GenAsyncCheckStringWithIndexCode(index)
		cCode, pCode := GenAsyncStringInputArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "boolean" {
		validateCode += validate.GenAsyncCheckBooleanWithIndexCode(index)
		code += GenAsyncBooleanInputArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "int" {
		validateCode += validate.GenAsyncCheckNumberWithIndexCode(index)
		code += GenAsyncIntInputArgTypeCode(arg.Name, sIndex, "int32")
	} else if arg.Type == "int32" {
		validateCode += validate.GenAsyncCheckNumberWithIndexCode(index)
		code += GenAsyncIntInputArgTypeCode(arg.Name, sIndex, "int32")
	} else if arg.Type == "int64" {
		validateCode += validate.GenAsyncCheckNumberWithIndexCode(index)
		code += GenAsyncIntInputArgTypeCode(arg.Name, sIndex, "int64")
	} else if arg.Type == "uint32" {
		validateCode += validate.GenAsyncCheckNumberWithIndexCode(index)
		code += GenAsyncIntInputArgTypeCode(arg.Name, sIndex, "unit32")
	} else if arg.Type == "float" {
		validateCode += validate.GenAsyncCheckNumberWithIndexCode(index)
		code += GenAsyncFloatInputArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "double" {
		code += GenAsyncDoubleInputArgTypeCode(arg.Name, sIndex)
	} else if arg.Type == "array" {
		validateCode += validate.GenAsyncCheckArrayWithIndexCode(index)
		cCode, pCode := GenAsyncArrayInputArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "object" {
		validateCode += validate.GenAsyncCheckObjectWithIndexCode(index)
		cCode, pCode := GenAsyncObjectInputArgTypeCode(arg.Name, sIndex, 2)
		code += cCode
		preCode += pCode
	} else if arg.Type == "arraybuffer" {
		validateCode += validate.GenAsyncCheckArrayBufferWithIndexCode(index)
		cCode, pCode := GenAsyncArrayBufferInputArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	} else if arg.Type == "callback" {
		validateCode += validate.GenAsyncCheckFunctionWithIndexCode(index)
		cCode, pCode := GenAsyncFuncitonInputArgTypeCode(arg.Name, sIndex)
		code += cCode
		preCode += pCode
	}

	code = genCheckInputArgCode(code, beforeCode, validateCode, index)
	return code, preCode
}
