package argsync

import (
	"fmt"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/validate"
)

func GenArgCode(args []config.Arg) (string, []string, string) {
	code := ""
	mCode := ""
	beforeCode := ""
	endCode := ""

	var argNames = make([]string, 0)
	for index, arg := range args {
		sIndex := fmt.Sprintf("%d", index)

		if arg.IsRequire {
			beforeCode += validate.GenRequireWithIndexCode(index)
		}

		if arg.Type == "int" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenIntArgTypeCode(arg.Name, sIndex, "int")
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "int32" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenIntArgTypeCode(arg.Name, sIndex, "int32")
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "int64" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenIntArgTypeCode(arg.Name, sIndex, "int64")
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "uint32" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenIntArgTypeCode(arg.Name, sIndex, "unit32")
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "float" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenFloatArgTypeCode(arg.Name, sIndex)
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "double" {
			code += validate.GenCheckNumberWithIndexCode(index)
			mCode += GenDoubleArgTypeCode(arg.Name, sIndex)
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "string" {
			code += validate.GenCheckStringWithIndexCode(index)
			sCode, eCode := GenStringArgTypeCode(arg.Name, sIndex)
			mCode += sCode
			endCode += eCode
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "boolean" {
			code += validate.GenCheckBooleanWithIndexCode(index)
			mCode += GenBooleanArgTypeCode(arg.Name, sIndex)
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "array" {
			code += validate.GenCheckArrayWithIndexCode(index)
			sCode, eCode := GenArrayArgTypeCode(arg.Name, sIndex)
			mCode += sCode
			endCode += eCode
			argNames = append(argNames, arg.Name)
		} else if arg.Type == "object" {
			code += validate.GenCheckObjectWithIndexCode(index)
			sCode, eCode := GenObjectArgTypeCode(arg.Name, sIndex)
			mCode += sCode
			endCode += eCode
			argNames = append(argNames, arg.Name)
		}
		/*
			 else if arg.Type == "arraybuffer" {
					code += validate.GenCheckArrayBufferWithIndexCode(index)
					sCode, eCode := GenArrayBufferArgTypeCode(arg.Name, sIndex)
					mCode += sCode
					endCode += eCode
					argNames = append(argNames, arg.Name)
				}
		*/
	}

	code = beforeCode + code + mCode
	return code, argNames, endCode

}
