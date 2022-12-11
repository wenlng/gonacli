package argsync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenArrayArgTypeCode(name string, index string) (string, string) {
	code := `
  Array wg__` + name + ` = Array::New(wg_env);
  if (wg_info.Length() > ` + index + `) {
    wg__` + name + ` = wg_info[` + index + `].As<Array>();
  }
  string wg_` + name + ` = wg_array_to_string(wg__` + name + `);
  char *` + name + ` = new char[wg_` + name + `.length() + 1];
  strcpy(` + name + `, wg_` + name + `.c_str());`

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
