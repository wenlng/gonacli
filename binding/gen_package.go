package binding

import (
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
)

func GenPackageFile(cfgs config.Config, packageName string) bool {
	code := `{
  "name": "` + cfgs.Name + `",
  "version": "1.0.0",
  "description": "",
  "main": "index.js",
  "scripts": {
    "install": "node-gyp rebuild",
    "build": "node-gyp configure && node-gyp build",
    "build:debug": "node-gyp configure && node-gyp build --debug",
    "build:release": "node-gyp configure && node-gyp build"
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {},
  "gypfile": true,
  "gyp": true
}
`
	// "bindings": "^1.5.0",
	// "node-addon-api": "^5.0.0"
	writePackageFile(code, packageName, cfgs.OutPut)
	return true
}

func writePackageFile(content string, filename string, outPath string) {
	outputDir := tools.FormatDirPath(outPath)
	tools.WriteFile(content, outputDir, filename)
}
