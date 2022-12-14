package content

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content/async"
	"github.com/wenlng/gonacli/content/base"
	"github.com/wenlng/gonacli/content/returns/resync"
	"github.com/wenlng/gonacli/tools"
)

func GenCode(config config.Config, cppName string) bool {
	code := base.GenHeaderFileCode(config.Name + ".h")

	curCode := ""
	asyncRegisterCode := ""
	hasAsync := false
	for _, export := range config.Exports {
		if export.JsCallMode == "async" {
			asyncHandlerCode, arCode := async.GenAsyncCallbackCode(export)
			curCode += asyncHandlerCode
			asyncRegisterCode += arCode
			hasAsync = true
		} else {
			if export.ReturnType == "int" {
				curCode += resync.GenReturnIntTypeCode(export, "int32")
			} else if export.ReturnType == "int32" {
				curCode += resync.GenReturnIntTypeCode(export, "int32")
			} else if export.ReturnType == "int64" {
				curCode += resync.GenReturnIntTypeCode(export, "int64")
			} else if export.ReturnType == "uint32" {
				curCode += resync.GenReturnIntTypeCode(export, "uint32")
			} else if export.ReturnType == "float" {
				curCode += resync.GenReturnFloatTypeCode(export)
			} else if export.ReturnType == "double" {
				curCode += resync.GenReturnDoubleTypeCode(export)
			} else if export.ReturnType == "boolean" {
				curCode += resync.GenReturnBooleanTypeCode(export)
			} else if export.ReturnType == "string" {
				curCode += resync.GenReturnStringTypeCode(export)
			} else if export.ReturnType == "array" {
				curCode += resync.GenReturnArrayTypeCode(export)
			} else if export.ReturnType == "object" {
				curCode += resync.GenReturnObjectTypeCode(export)
			}
			/*
				else if export.ReturnType == "arraybuffer" {
					curCode += resync.GenReturnArrayBufferTypeCode(export)
				}
			*/
		}
	}

	code += base.GenBeforeCode(hasAsync)
	code += curCode
	code += base.GenAfterCode(config, asyncRegisterCode)

	writeFile(code, cppName, config.OutPut)
	return true
}

func writeFile(content string, filename string, outPath string) {
	outputDir := tools.FormatDirPath(outPath)
	tools.WriteFile(content, outputDir, filename)
}
