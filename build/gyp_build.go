package build

import (
	"node_addon_go/check"
	"node_addon_go/clog"
	"node_addon_go/cmd"
	"node_addon_go/config"
	"node_addon_go/tools"
	"path/filepath"
)

func MakeToAddon(cfgs config.Config, args string, makeMpn bool) bool {
	done := true

	path := tools.FormatDirPath(cfgs.OutPut)

	// 检查配置文件
	if err := check.CheckBaseConfig(cfgs); err != nil {
		clog.Error(err)
		done = false
		return done
	}
	if c := check.CheckExportApiWithSourceFile(cfgs); !c {
		done = false
	}

	// 清空生成的相关文件
	_ = tools.RemoveDirContents(filepath.Join(path, "build"))
	files := []string{
		filepath.Join(path, "package-lock.json"),
	}
	_ = tools.RemoveFiles(files)

	if makeMpn {
		clog.Info("Staring install npm dependencies ...")
		msg, err := cmd.RunCommand(
			"./",
			"cd "+path+" && npm install",
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

	return done
}
