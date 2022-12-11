package argsync

import (
	"node_addon_go/tools"
)

func GenStringArgTypeCode(name string, index string) (string, string) {
	code := `
  string wg_` + name + ` = "";
  if (wg_info.Length() > ` + index + `) {
    wg_` + name + ` = wg_info[` + index + `].As<String>().Utf8Value();
  }
  char *` + name + ` = new char[wg_` + name + `.length() + 1];
  strcpy(` + name + `, wg_` + name + `.c_str());`

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
