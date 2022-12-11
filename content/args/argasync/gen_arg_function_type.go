package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncFuncitonArgTypeCode(name string, index string) (string, string) {
	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  char * ` + name + ` = new char[wg_` + name + `Info->len];
  strcpy(` + name + `, (char *)wg_` + name + `Info->value);
  // ---- `

	//tools.FormatCodeIndentLn(`char * `+name+` = "`+name+`(){}";`, 2)
	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}

func GenAsyncFuncitonInputArgTypeCode(name string, index string) (string, string) {
	/*
		string __` + name + ` = "` + name + `(){}";
	*/
	code := `
  // arg - ` + name + `
  string wg__` + name + ` = "` + name + `(){}";
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsFunction()){
    wg__` + name + ` = wg_v` + index + `.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * ` + name + ` = new char[wg__` + name + `.length() + 1];
  strcpy(` + name + `, wg__` + name + `.c_str());
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=1;
  wg_addon->args[` + index + `]->len=wg__` + name + `.length() + 1;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	//tools.FormatCodeIndentLn(`char * `+name+` = "`+name+`(){}";`, 2)
	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
