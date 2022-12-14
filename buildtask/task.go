package buildtask

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/wenlng/gonacli/buildtask/compatible"
	"github.com/wenlng/gonacli/check"
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

func checkConfigure(c config.Config) bool {
	// 检查配置文件
	if err := check.CheckBaseConfig(c); err != nil {
		clog.Error(err)
		return false
	}
	if err := check.CheckAsyncCorrectnessConfig(c); err != nil {
		clog.Error(err)
		return false
	}
	if c := check.CheckExportApiWithSourceFile(c); !c {
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

	// 检测配置文件
	if c := checkConfigure(cfgs); !c {
		return
	}

	if len(args) > 0 {
		// 删除 args 两边 ' 或 "
		if strings.Index(args, "'") == 0 || strings.Index(args, "\"") == 0 {
			args = args[1:]
			if strings.LastIndex(args, "'") == len(args)-1 {
				args = args[:len(args)-1]
			}
		} else if strings.Index(args, "\"") == 0 {
			args = args[1:]
			if strings.LastIndex(args, "\"") == len(args)-1 {
				args = args[:len(args)-1]
			}
		}
	}

	// 生成 golang lib
	if d := buildGoToLibrary(cfgs, args); !d {
		clog.Error("Fail build golang lib!")
		return
	}

	clog.Success("Successfully build golang lib ~")
	fmt.Println("")
}

// 生成 bridge c/c++ 代码
func RunGenerateTask(config string) {
	if ok := parseConfig(config); !ok {
		return
	}

	// 检测配置文件
	if c := checkConfigure(cfgs); !c {
		return
	}

	if done := generateAddonBridge(cfgs); !done {
		clog.Error("Fail generated bridge code!")
		return
	}

	clog.Success("Successfully generate the Addon bridge c/c++ code of Nodejs ~")
	//clog.Info("Please execute the following command to make the addon of nodejs ~")
	//clog.Info(fmt.Sprintf("> cd %s && gonacli make --config xxx.json", cfgs.OutPut))
	//clog.Info(fmt.Sprintf("or > cd %s && node-gyp configure && node-gyp build", cfgs.OutPut))
	//clog.Info(fmt.Sprintf("or > cd %s && npm install && npm run build", cfgs.OutPut))
	fmt.Println("")
}

// 初始化 npm install xxxx
func RunInstallTask(config string) {
	if ok := parseConfig(config); !ok {
		return
	}

	// 检测配置文件
	if c := checkConfigure(cfgs); !c {
		return
	}

	// 生成扩展
	if done := installDep(cfgs); !done {
		clog.Error("Fail installed!")
		return
	}

	clog.Success("Successfully installed ~")
	fmt.Println("")
}

// 编译 node addon
func RunMakeTask(config string, args string) {
	if ok := parseConfig(config); !ok {
		return
	}

	// 检测配置文件
	if c := checkConfigure(cfgs); !c {
		return
	}

	if len(args) > 0 {
		// 删除 args 两边 ' 或 "
		if strings.Index(args, "'") == 0 || strings.Index(args, "\"") == 0 {
			args = args[1:]
			if strings.LastIndex(args, "'") == len(args)-1 {
				args = args[:len(args)-1]
			}
		} else if strings.Index(args, "\"") == 0 {
			args = args[1:]
			if strings.LastIndex(args, "\"") == len(args)-1 {
				args = args[:len(args)-1]
			}
		}
	}

	// 生成扩展
	if done := makeToAddon(cfgs, args); !done {
		clog.Error("Fail make addon!")
		return
	}

	clog.Success("Successfully make the addon of Nodejs ~")
	fmt.Println("")
}

// window 环境下兼容处理
func RunMsvcTask(config string, useVS bool, msvc32Vs bool) {
	if ok := parseConfig(config); !ok {
		return
	}

	// 检测配置文件
	if c := checkConfigure(cfgs); !c {
		return
	}

	// 修复兼容 MSVC
	clog.Info("Staring fix file ...")
	if tools.IsWindowsOs() {
		compatible.FixCGOWithWindow(cfgs)
	}
	clog.Info("Fix file done ~")

	// 生成dll
	clog.Info("Staring build dll ...")
	if done := buildToDll(cfgs); !done {
		clog.Error("Fail build dll!")
		return
	}
	clog.Success("Successfully build dll ~")
	fmt.Println("")

	// 生成 lib
	clog.Info("Staring build lib ...")
	if done := buildToMSVCLib(cfgs, useVS, msvc32Vs); !done {
		clog.Error("Fail build lib!")
		return
	}

	clog.Success("Successfully build lib ~")
	fmt.Println("")
}
