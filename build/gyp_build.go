package build

import (
	"github.com/wenlng/gonacli/check"
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/cmd"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"path/filepath"
)

func MakeToAddon(cfgs config.Config, args string, makeMpn bool) bool {

	path := tools.FormatDirPath(cfgs.OutPut)

	// 检查配置文件
	if err := check.CheckBaseConfig(cfgs); err != nil {
		clog.Error(err)
		return false
	}
	if c := check.CheckExportApiWithSourceFile(cfgs); !c {
		return false
	}

	// 检测是否存在 node_modules 目录
	if !tools.Exists(filepath.Join(path, "node_modules")) && !makeMpn {
		clog.Error("When making for the first time, you need to add \"--npm-i\" args install dependencies.")
		return false
	}

	// 清空生成的相关文件
	_ = tools.RemoveDirContents(filepath.Join(path, "build"))
	files := []string{
		filepath.Join(path, "package-lock.json"),
	}
	_ = tools.RemoveFiles(files)

	if makeMpn {
		clog.Info("Staring install npm dependencies ...")
		// "bindings": "^1.5.0",
		// "node-addon-api": "^5.0.0"
		msg, err := cmd.RunCommand(
			"./",
			"cd "+path+" && npm install bindings node-addon-api -S",
		)
		if err != nil {
			clog.Error(err)
			return false
		}
		clog.Info(msg)
		clog.Info("Install npm dependencies done ~")
	}

	clog.Info("Staring make addon ...")
	msg, err := cmd.RunCommand(
		"./",
		"cd "+path+" && node-gyp configure && node-gyp build "+args,
	)
	if err != nil {
		clog.Error(err)
		return false
	}
	clog.Info("Make addon done ~")
	clog.Info(msg)

	return true
}
