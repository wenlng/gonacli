package binding

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
)

func GenGypFile(cfgs config.Config, bindingName string) bool {
	code := `{
    "targets": [
        {
            "target_name": "` + cfgs.Name + `",
            "cflags!": [ "-fno-exceptions" ],
            "cflags_cc!": [ "-fno-exceptions" ],
            "sources": [ "` + cfgs.Name + `.cc" ],
            "include_dirs": [
                "<!@(node -p \"require('node-addon-api').include\")"
            ],
            'defines': [ 'NAPI_DISABLE_CPP_EXCEPTIONS' ],
            "libraries": [
                "../` + cfgs.Name + `.a"
            ],
        }
    ]
}`

	writeGypFile(code, bindingName, cfgs.OutPut)
	return true
}

func writeGypFile(content string, filename string, outPath string) {
	outputDir := tools.FormatDirPath(outPath)
	tools.WriteFile(content, outputDir, filename)
}
