package argasync

func GenAsyncFloatArgTypeCode(name string, index string) string {
	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  float ` + name + ` = (float)wg_` + name + `Info->value;
  // ---- `
	return code
}

func GenAsyncFloatInputArgTypeCode(name string, index string) string {
	/*
		napi_status sts` + index + ` = napi_get_value_double(env, addon->args[` + index + `], &` + name + `);
		  wg_catch_err(env, sts` + index + `);
	*/
	code := `
  // arg - ` + name + `
  float ` + name + ` = 0.0;
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsNumber()){
    ` + name + ` =	wg_v` + index + `.As<Number>().FloatValue();
  }
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=3;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `

	return code
}
