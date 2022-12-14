package binding

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenJsCallIndexFile(cfgs config.Config, indexJsName string) bool {
	apiCode := genJsCallApiListCode(cfgs.Exports, cfgs.Name)
	code := `const ` + cfgs.Name + ` = require('bindings')('` + cfgs.Name + `');
// JS call API
module.exports = { ` + apiCode + `
};
`
	writeJsFile(code, indexJsName, cfgs.OutPut)
	return true
}

func genJsCallApiListCode(exports []config.Export, name string) string {
	codeList := make([]string, 0)

	for _, export := range exports {
		jsCallApiName := export.JsCallName
		jsCallExportName := tools.ToFirstCharLower(jsCallApiName)
		codeList = append(codeList, tools.FormatCodeIndentLn(jsCallExportName+" : "+name+"."+jsCallApiName, 2))
	}

	return strings.Join(codeList, ",")
}

func writeJsFile(content string, filename string, outPath string) {
	outputDir := tools.FormatDirPath(outPath)
	tools.WriteFile(content, outputDir, filename)
}
