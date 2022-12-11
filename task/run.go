package task

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/wenlng/gonacli/build"
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"path/filepath"
	"strings"
)

var cfgs config.Config

func parseConfig(config string) bool {
	configPath := "./goaddon.json"
	if len(config) > 0 {
		configPath = config
	}

	cpath := tools.FormatDirPath(configPath)

	if !tools.Exists(configPath) {
		clog.Error("The json configuration file of addon does not exist.")
		return false
	}

	path := filepath.Join(cpath)
	if err := configor.Load(&cfgs, path); err != nil {
		clog.Error(err)
		return false
	}

	return true
}

// 编译 golang lib 文件
// gonacli build => go build -buildmode c-archive -o xxx.a xxx.go xxx1.go xxx2.go ...
func RunBuildTask(config string, args string) {
	if ok := parseConfig(config); !ok {
		return
	}

	// 删除 args 两边 ' 或 "
	if strings.Index(args, "'") == 0 || strings.Index(args, "\"") == 0 {
		args = args[1:]
	}

	if strings.LastIndex(args, "'") == len(args)-1 || strings.LastIndex(args, "\"") == len(args)-1 {
		args = args[:len(args)-1]
	}

	// 生成 golang lib
	done := true
	if d := build.BuildGoToLibrary(cfgs, args); !d {
		done = false
	}

	if done {
		buildDon(cfgs)
	} else {
		buildFail(cfgs)
	}
}
func buildFail(cfgs config.Config) {
	clog.Error("Fail build golang lib!")
}
func buildDon(cfgs config.Config) {
	clog.Success("Successfully build golang lib ~")
	fmt.Println("")
}

// 生成 bridge c/c++ 代码
func RunGenerateTask(config string) {
	if ok := parseConfig(config); !ok {
		return
	}

	if done := build.GenerateAddonBridge(cfgs); done {
		generateDon(cfgs)
	} else {
		generateFail(cfgs)
	}
}
func generateFail(cfgs config.Config) {
	clog.Error("Fail generated bridge code!")
}
func generateDon(cfgs config.Config) {
	clog.Success("Successfully generate the Addon bridge c/c++ code of Nodejs ~")
	clog.Info("Please execute the following command to make the addon of nodejs ~")
	clog.Info(fmt.Sprintf("> cd %s && gonacli make --config xxx.json", cfgs.OutPut))
	clog.Info(fmt.Sprintf("or > cd %s && node-gyp configure && node-gyp build", cfgs.OutPut))
	clog.Info(fmt.Sprintf("or > cd %s && npm install && npm run build", cfgs.OutPut))
	fmt.Println("")
}

// 编译 node addon 扩展
func RunMakeTask(config string, args string, makeMpn bool) {
	if ok := parseConfig(config); !ok {
		return
	}

	if done := build.MakeToAddon(cfgs, args, makeMpn); done {
		nodeMakeDon(cfgs)
	} else {
		nodeMakeFail(cfgs)
	}
}
func nodeMakeFail(cfgs config.Config) {
	clog.Error("Fail make addon!")
}
func nodeMakeDon(cfgs config.Config) {
	clog.Success("Successfully make the addon of Nodejs ~")
	fmt.Println("")
}
