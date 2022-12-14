package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncArrayBufferArgTypeCode(name string, index string) (string, string) {
	//char *` + name + ` = new char[wg_` + name + `Info->len];
	//strcpy(` + name + `, (char *)wg_` + name + `Info->value);

	code := `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  GoSlice ` + name + ` = wg_build_go_slice(wg_` + name + `Info->value, wg_` + name + `Info->len, wg_` + name + `Info->len);`

	//endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	endCode := ""
	return code, endCode
}

func GenAsyncArrayBufferInputArgTypeCode(name string, index string) (string, string) {
	/*
	 Array(env, addon->args[` + index + `]);
	*/
	code := `
  // arg - ` + name + `
  ArrayBuffer wg_` + name + ` = ArrayBuffer();
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsArrayBuffer()){
    wg_` + name + ` = wg_v` + index + `.As<ArrayBuffer>();
  }
  char *` + name + ` = (char *)wg_` + name + `.Data();
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=1;
  wg_addon->args[` + index + `]->len=strlen(` + name + `);
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
