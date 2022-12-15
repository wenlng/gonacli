package buildtask

import (
	"github.com/wenlng/gonacli/binding"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/content"
	"github.com/wenlng/gonacli/tools"
	"path/filepath"
)

func generateAddonBridge(cfgs config.Config) bool {

	cppName := cfgs.Name + ".cc"
	bindingName := "binding.gyp"
	indexJsName := "index.js"
	indexDTsName := "index.d.ts"
	packageName := "package.json"

	// 清空生成的相关文件
	outputDir := tools.FormatDirPath(cfgs.OutPut)
	paths := []string{
		filepath.Join(outputDir, cppName),
		filepath.Join(outputDir, bindingName),
		filepath.Join(outputDir, indexJsName),
		filepath.Join(outputDir, indexDTsName),
		filepath.Join(outputDir, packageName),
	}
	//_ = tools.RemoveDirContents(outputDir)
	_ = tools.RemoveFiles(paths)

	// 生成 addon c/c++ 代码
	if g := content.GenCode(cfgs, cppName); !g {
		//clog.Warning("Please check whether the \"goaddon\" configuration file is correct.")
		return false
	}

	// 生成 node-gyp 编译配置文件
	if y := binding.GenGypFile(cfgs, bindingName); !y {
		return false
	}

	// 生成 js call api to index.js
	if i := binding.GenJsCallIndexFile(cfgs, indexJsName); !i {
		return false
	}

	// 生成 js call api to index.d.t
	if t := binding.GenJsCallDeclareIndexFile(cfgs, indexDTsName); !t {
		return false
	}

	// 生成 npm package 包模板文件
	if p := binding.GenPackageFile(cfgs, packageName); !p {
		return false
	}

	return true
}
