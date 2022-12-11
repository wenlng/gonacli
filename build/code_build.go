package build

import (
	binding2 "node_addon_go/binding"
	"node_addon_go/check"
	"node_addon_go/clog"
	"node_addon_go/config"
	"node_addon_go/content"
	"node_addon_go/tools"
	"path/filepath"
)

func GenerateAddonBridge(cfgs config.Config) bool {
	done := true
	// 检查配置文件
	if err := check.CheckBaseConfig(cfgs); err != nil {
		clog.Error(err)
		done = false
		return done
	}
	if c := check.CheckExportApiWithSourceFile(cfgs); !c {
		done = false
	}

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
		done = false
	}

	// 生成 node-gyp 编译配置文件
	if y := binding2.GenGypFile(cfgs, bindingName); !y {
		done = false
	}

	// 生成 js call api to index.js
	if i := binding2.GenJsCallIndexFile(cfgs, indexJsName); !i {
		done = false
	}

	// 生成 js call api to index.d.t
	if t := binding2.GenJsCallDeclareIndexFile(cfgs, indexDTsName); !t {
		done = false
	}

	// 生成 npm package 包模板文件
	if p := binding2.GenPackageFile(cfgs, packageName); !p {
		done = false
	}

	return done
}
