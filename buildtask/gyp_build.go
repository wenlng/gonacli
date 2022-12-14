package buildtask

import (
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/cmd"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"path/filepath"
)

func makeToAddon(cfgs config.Config, args string) bool {

	path := tools.FormatDirPath(cfgs.OutPut)

	// 检测是否存在执行过 gonacli generate 命令
	if !tools.Exists(filepath.Join(path, cfgs.Name+".cc")) {
		clog.Error("You need to run \"gonacli generate\" generate c/c++ bridge code.")
		return false
	}

	// 检测是否存在执行过 gonacli build 命令
	if !tools.Exists(filepath.Join(path, cfgs.Name+".a")) {
		clog.Error("You need to run \"gonacli build\" build golang lib.")
		return false
	}

	// 检测是否存在执行过 gonacli install 命令
	if !tools.Exists(filepath.Join(path, "node_modules")) {
		clog.Error("You need to run \"gonacli install\" install dependencies.")
		return false
	}

	// window 环境检测是否存在执行过 gonacli msvc 命令
	if tools.IsWindowsOs() {
		if !tools.Exists(filepath.Join(path, cfgs.Name+".lib")) {
			clog.Error("You need to run \"gonacli msvc\" build lib on windows OS.")
			return false
		}
	}

	// 清空生成的相关文件
	_ = tools.RemoveDirContents(filepath.Join(path, "build"))
	files := []string{
		filepath.Join(path, "package-lock.json"),
	}
	_ = tools.RemoveFiles(files)

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
