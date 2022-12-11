package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncArrayArgTypeCode(name string, index string) (string, string) {
	//   Array wg__` + name + ` = Array(wg_env, wg_addon->args[` + index + `]);
	code := `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  char *` + name + ` = new char[wg_` + name + `Info->len];
  strcpy(` + name + `, (char *)wg_` + name + `Info->value);`

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}

func GenAsyncArrayInputArgTypeCode(name string, index string) (string, string) {
	/*
	 Array(env, addon->args[` + index + `]);
	*/
	code := `
  // arg - ` + name + `
  Array wg__` + name + ` = Array::New(wg_env);
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsArray()){
    wg__` + name + ` = wg_v` + index + `.As<Array>();
  }
  string wg_` + name + ` = wg_array_to_string(wg__` + name + `);
  char *` + name + ` = new char[wg_` + name + `.length() + 1];
  strcpy(` + name + `, wg_` + name + `.c_str());
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=1;
  wg_addon->args[` + index + `]->len=wg_` + name + `.length() + 1;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
