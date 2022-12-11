package argasync

import (
	"github.com/wenlng/gonacli/tools"
)

func GenAsyncArrayBufferArgTypeCode(name string, index string) (string, string) {
	//code := `
	//byte * wg_` + name + ` = (byte *)wg_addon->args[` + index + `];
	//char * ` + name + ` = (char *)wg_` + name + `;`

	code := `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  char *` + name + ` = new char[wg_` + name + `Info->len];
  strcpy(` + name + `, (char *)wg_` + name + `Info->value);`

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
  byte *` + name + ` = (byte *)wg_` + name + `.Data();
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=6;
  wg_addon->args[` + index + `]->len=strlen((char *)` + name + `);
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	endCode := tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	return code, endCode
}
