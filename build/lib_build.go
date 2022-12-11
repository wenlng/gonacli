package build

import (
	"node_addon_go/check"
	"node_addon_go/clog"
	"node_addon_go/cmd"
	"node_addon_go/config"
	"node_addon_go/tools"
	"path/filepath"
	"strings"
)

func BuildGoToLibrary(cfgs config.Config, args string) bool {

	done := true
	// 检查配置文件
	if err := check.CheckBaseConfig(cfgs); err != nil {
		clog.Error(err)
		done = false
	}
	if err := check.CheckAsyncCorrectnessConfig(cfgs); err != nil {
		clog.Error(err)
		done = false
	}
	if c := check.CheckExportApiWithSourceFile(cfgs); !c {
		done = false
	}

	libName := cfgs.Name + ".a"
	libHName := cfgs.Name + ".h"

	// 清空生成的相关文件
	outputDir := tools.FormatDirPath(cfgs.OutPut)
	paths := []string{
		filepath.Join(outputDir, libName),
		filepath.Join(outputDir, libHName),
	}
	_ = tools.RemoveFiles(paths)

	sourceFiles := genBuildFile(cfgs)
	if d := buildGoToLibrary(sourceFiles, libName, outputDir, args); !d {
		done = false
	}

	return done
}

// 生成 build go 文件集合
func genBuildFile(config config.Config) string {
	files := ""
	for _, source := range config.Sources {
		files += " " + source
	}
	return files
}

func buildGoToLibrary(sourceFiles string, libName string, outPath string, args string) bool {
	msg, err := cmd.RunCommand(
		"./",
		"go build -buildmode c-archive "+args+" -o "+formatLibOutPath(libName, outPath)+" "+sourceFiles,
	)
	if err != nil {
		clog.Error(err)
		return false
	}
	clog.Info(msg)
	return true
}

func formatLibOutPath(libName string, outPath string) string {
	if strings.LastIndex(outPath, "/") == len(outPath)-1 {
		return outPath + libName
	}
	return outPath + "/" + libName
}
