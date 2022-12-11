package argasync

func GenAsyncDoubleArgTypeCode(name string, index string) string {
	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  double ` + name + ` = (double)wg_` + name + `Info->value;
  // ---- `
	return code
}

func GenAsyncDoubleInputArgTypeCode(name string, index string) string {
	/*
	  napi_status sts` + index + ` = napi_get_value_double(env, addon->args[` + index + `], &` + name + `);
	  wg_catch_err(env, sts` + index + `);
	*/
	code := `
  // arg - ` + name + `
  double ` + name + `;
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsNumber()){
    ` + name + ` =	wg_v` + index + `.As<Number>().DoubleValue();
  }
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=4;
  wg_addon->args[` + index + `]->value=(void *)` + name + `;
  // ---- `
	return code
}
