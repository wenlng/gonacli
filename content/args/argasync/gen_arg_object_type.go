package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncObjectArgTypeCode(name string, index string, v int) (string, string) {
	//code := `
	//Object wg__` + name + ` = Object(wg_env, wg_addon->args[` + index + `]);
	//string wg_` + name + ` = wg_object_to_string(wg__` + name + `);
	//char *` + name + ` = new char[wg_` + name + `.length() + 1];
	//strcpy(` + name + `, wg_` + name + `.c_str());`
	code := `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  char *` + name + ` = new char[wg_` + name + `Info->len];
  strcpy(` + name + `, (char *)wg_` + name + `Info->value);`

	endCode := tools.FormatCodeIndent(`delete [] `+name+`;`, 2)
	return code, endCode
}

func GenAsyncObjectInputArgTypeCode(name string, index string, v int) (string, string) {
	/*
	  Object __` + name + ` = Object(env, addon->args[` + index + `]);
	*/
	code := `
  // arg - ` + name + `
  Object wg__` + name + ` = Object::New(wg_env);
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsObject()){
    wg__` + name + ` = wg_v` + index + `.As<Object>();
  }
  string wg_` + name + ` = wg_object_to_string(wg__` + name + `);
  char *` + name + ` = new char[wg_` + name + `.length() + 1];
  strcpy(` + name + `, wg_` + name + `.c_str());
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=1;
  wg_addon->args[` + index + `]->len=wg_` + name + `.length() + 1;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	endCode := tools.FormatCodeIndent(`delete [] `+name+`;`, 2)
	return code, endCode
}
