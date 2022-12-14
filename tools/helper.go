package tools

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

func GetPWD() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

func EnsureDir(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return
		}
	}
	return
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func WriteFile(content string, dirname string, filename string) error {
	_ = os.MkdirAll(dirname, 0777)
	file := dirname + filename
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Println(">>> ", err)
		return err
	}

	defer logFile.Close()
	_, err1 := io.WriteString(logFile, content)
	if err1 != nil {
		log.Println(">>> ", err1)
		return err
	}
	return nil
}

func CheckOS64Unit() bool {
	unit := 32 << (^uint(0) >> 63)
	return unit >= 64
}

func FormatCodeIndent(str string, indent int) string {
	newStr := ""
	for i := 0; i < indent; i++ {
		newStr += ` `
	}

	newStr += str
	return newStr
}

func FormatCodeIndentLn(str string, indent int) string {
	newStr := `
`
	for i := 0; i < indent; i++ {
		newStr += ` `
	}

	newStr += str
	return newStr
}

func InSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func IndexSlice(items []string, item string) int {
	index := -1
	for i, eachItem := range items {
		if eachItem == item {
			index = i
			return index
		}
	}
	return index
}

func IsWindowsOs() bool {
	return runtime.GOOS == "windows"
}

func IsLinuxOs() bool {
	return runtime.GOOS == "linux"
}

func FormatDirPath(op string) string {
	if strings.LastIndex(op, "/") != len(op)-1 {
		op += "/"
	}
	if strings.Index(op, "/") == 0 {
		return op
	}
	return GetPWD() + "/" + op
}

func ToFirstLower(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return str
}

func ToFirstUpper(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return str
}

func ToFirstCharLower(s string) string {
	arr := strings.Split(s, "_")
	newStr := ""
	for i, s2 := range arr {
		if i == 0 {
			newStr += ToFirstLower(s2)
		} else {
			newStr += ToFirstUpper(s2)
		}
	}
	return newStr
}

func RemoveDirContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}

	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveFiles(files []string) error {
	for _, file := range files {
		if Exists(file) {
			err := os.RemoveAll(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func RenameFile(spath string, tpath string) error {
	return os.Rename(spath, tpath)
}
