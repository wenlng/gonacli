package binding

import (
	"node_addon_go/config"
	"node_addon_go/tools"
	"strings"
)

func GenJsCallDeclareIndexFile(cfgs config.Config, indexDTsName string) bool {
	code := "// This is binding js call api declare"
	code += genJsCallApiListDeclareCode(cfgs.Exports, cfgs.Name)

	writeJsDeclareFile(code, indexDTsName, cfgs.OutPut)
	return true
}

var allowArgsList = []string{"int", "int32", "int64", "uint32", "float", "double", "boolean", "string", "array", "object", "arraybuffer", "callback"}
var mapDeclareList = []string{"number", "number", "number", "number", "number", "number", "boolean", "string", "Array<any>", "object", "ArrayBuffer", "any"}

func genJsCallApiListDeclareCode(exports []config.Export, name string) string {
	codeList := make([]string, 0)

	for _, export := range exports {
		jsCallApiName := export.JsCallName
		jsCallExportName := tools.ToFirstCharLower(jsCallApiName)

		rtIndex := tools.IndexSlice(allowArgsList, export.ReturnType)
		declareType := mapDeclareList[rtIndex]
		argCode := genArgTypeList(export.Args)

		codeList = append(codeList, tools.FormatCodeIndentLn(`export declare function `+jsCallExportName+`(`+argCode+`): `+declareType+`;`, 0))
	}

	return strings.Join(codeList, "")
}

func genArgTypeList(args []config.Arg) string {
	// source: object, config: object
	codeList := make([]string, 0)

	lastRequireIndex := 0
	for i, arg := range args {
		if arg.IsRequire {
			lastRequireIndex = i
		}
	}

	for i, arg := range args {
		rtIndex := tools.IndexSlice(allowArgsList, arg.Type)
		declareType := mapDeclareList[rtIndex]

		if !arg.IsRequire && lastRequireIndex < i {
			codeList = append(codeList, arg.Name+" ?: "+declareType)
		} else {
			codeList = append(codeList, arg.Name+" : "+declareType)
		}
	}
	return strings.Join(codeList, ", ")
}

func writeJsDeclareFile(content string, filename string, outPath string) {
	outputDir := tools.FormatDirPath(outPath)
	tools.WriteFile(content, outputDir, filename)
}
