package buildtask

import (
	"github.com/wenlng/gonacli/clog"
	"github.com/wenlng/gonacli/cmd"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"path/filepath"
)

func buildGoToLibrary(cfgs config.Config, args string) bool {
	libName := cfgs.Name + ".a"
	libHName := cfgs.Name + ".h"

	// 清空生成的相关文件
	outputDir := tools.FormatDirPath(cfgs.OutPut)
	paths := []string{
		filepath.Join(outputDir, libName),
		filepath.Join(outputDir, libHName),
	}
	_ = tools.RemoveFiles(paths)

	clog.Info("Start build library ...")
	sourceFiles := genBuildFile(cfgs)
	if d := buildLibrary(sourceFiles, libName, outputDir, args); !d {
		return false
	}

	return true
}

// 生成 build go 文件集合
func genBuildFile(config config.Config) string {
	files := ""
	for _, source := range config.Sources {
		files += " " + source
	}
	return files
}

func buildLibrary(sourceFiles string, libName string, outPath string, args string) bool {
	oPath := outPath + libName
	msg, err := cmd.RunCommand(
		"./",
		"go build -buildmode c-archive "+args+" -o "+oPath+sourceFiles,
	)
	if err != nil {
		clog.Error(err)
		return false
	}
	clog.Info(msg)
	return true
}
