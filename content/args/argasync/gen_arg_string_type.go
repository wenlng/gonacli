package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncStringArgTypeCode(name string, index string) (string, string) {
	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  char * ` + name + ` = new char[wg_` + name + `Info->len];
  strcpy(` + name + `, (char *)wg_` + name + `Info->value);
  // ---- `

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}

func GenAsyncStringInputArgTypeCode(name string, index string) (string, string) {
	/*
			napi_status sts` + index + ` = napi_get_value_string_utf8(env, args[` + index + `], &_` + name + `[0], _` + name + `.capacity(), nullptr);
		      wg_catch_err(env, sts` + index + `);
	*/
	code := `
  // arg - ` + name + `
  string wg_` + name + ` = "";
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsString()){
    wg_` + name + ` = wg_v` + index + `.As<String>().Utf8Value();
  }
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
