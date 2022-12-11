package argasync

func GenAsyncBooleanArgTypeCode(name string, index string) string {
	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  bool ` + name + ` = (bool)wg_` + name + `Info->value;
  // ---- `
	return code
}

func GenAsyncBooleanInputArgTypeCode(name string, index string) string {
	/*
		napi_status sts` + index + ` = napi_get_value_bool(env, args[` + index + `], &` + name + `);
		    wg_catch_err(env, sts` + index + `);
	*/
	code := `
  // arg - ` + name + `
  bool ` + name + ` = false;
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsBoolean()){
    ` + name + ` =	wg_v` + index + `.As<Boolean>();
  }
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=5;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	return code
}
