package compatible

import (
	"bufio"
	"github.com/wenlng/gonacli/config"
	"github.com/wenlng/gonacli/tools"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func FixCGOWithWindow(cfgs config.Config) bool {
	outputDir := tools.FormatDirPath(cfgs.OutPut)
	headerFileName := cfgs.Name + ".h"
	bakHeaderFileName := cfgs.Name + ".s.h"
	headerFilePath := filepath.Join(outputDir, headerFileName)
	bakHeaderFilePath := filepath.Join(outputDir, bakHeaderFileName)
	paths := []string{
		bakHeaderFilePath,
	}
	_ = tools.RemoveFiles(paths)

	// 重命名
	_ = tools.RenameFile(headerFilePath, bakHeaderFilePath)

	return FixCGOHeaderFile(bakHeaderFilePath, outputDir, headerFileName)
}

func FixCGOHeaderFile(inputFilePath string, outPath string, filename string) bool {
	fp, err := os.OpenFile(inputFilePath, os.O_RDONLY, 0666)
	if err != nil {
		return false
	}

	defer fp.Close()
	reader := bufio.NewReader(fp)

	results := ""
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		content := string(line)
		content = fixCGOHeaderCode(content)
		if len(results) <= 0 {
			results = tools.FormatCodeIndentLn(content, 0)
		} else {
			results += tools.FormatCodeIndentLn(content, 0)
		}
	}

	tools.WriteFile(results, outPath, filename)
	return true
}

func fixCGOHeaderCode(content string) string {
	// remove
	// #line 1 "cgo-builtin-export-prolog"
	// #line 1 "cgo-gcc-export-header-prolog"
	// typedef __SIZE_TYPE__ GoUintptr;
	// typedef float _Complex GoComplex64;
	// typedef double _Complex GoComplex128;

	if strings.Index(content, "#line 1") == 0 && strings.Contains(content, "cgo-builtin-export-prolog") {
		content = "/*" + content + "*/"
	}

	if strings.Index(content, "#line 1") == 0 && strings.Contains(content, "cgo-gcc-export-header-prolog") {
		content = "/*" + content + "*/"
	}

	if strings.Index(content, "typedef") == 0 && strings.Contains(content, "GoUintptr") {
		content = "/*" + content + "*/"
	}
	if strings.Index(content, "typedef") == 0 && strings.Contains(content, "GoComplex64") {
		content = "/*" + content + "*/"
	}
	if strings.Index(content, "typedef") == 0 && strings.Contains(content, "GoComplex128") {
		content = "/*" + content + "*/"
	}

	return content
}
