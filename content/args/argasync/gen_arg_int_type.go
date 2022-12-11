package argasync

func GenAsyncIntArgTypeCode(name string, index string, varType string) string {
	// 	int64 uint32 int32
	//  int64_t uint32_t int32_t
	// napi_get_value_int32
	// napi_get_value_uint32
	// napi_get_value_int64

	typeStr := "int"
	if varType == "int64" {
		typeStr = "long long int"
	} else if varType == "uint32" {
		typeStr = "unsigned long"
	}

	code := `
  // arg - ` + name + `
  WgAddonArgInfo * wg_` + name + `Info = wg_addon->args[` + index + `];
  ` + typeStr + ` ` + name + ` = (` + typeStr + `)(intptr_t)wg_` + name + `Info->value;
  // ---- `
	return code
}

func GenAsyncIntInputArgTypeCode(name string, index string, varType string) string {
	// 	int64 uint32 int32
	//  int64_t uint32_t int32_t
	// napi_get_value_int32
	// napi_get_value_uint32
	// napi_get_value_int64

	typeStr := "int"
	//getTypeStr := "napi_get_value_int32"
	getValue := name + ` = wg_v` + index + `.As<Number>().Int32Value();`
	if varType == "int64" {
		typeStr = "long long int "
		//getTypeStr = "napi_get_value_int64"
		getValue = name + ` = wg_v` + index + `.As<Number>().Int64Value();`
	} else if varType == "uint32" {
		typeStr = "unsigned long"
		//getTypeStr = "napi_get_value_uint32"
		getValue = name + ` = wg_v` + index + `.As<Number>().Uint32Value();`
	}

	/*
	  napi_status sts` + index + ` = ` + getTypeStr + `(env, args[` + index + `], &` + name + `)
	  wg_catch_err(env, sts` + index + `);
	*/
	code := `
  // arg - ` + name + `
  ` + typeStr + ` ` + name + ` = 0;
  if(!wg_v` + index + `.IsUndefined() && wg_v` + index + `.IsNumber()){
    ` + getValue + `
  }
  wg_addon->args[` + index + `] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[` + index + `]));
  wg_addon->args[` + index + `]->type=2;
  wg_addon->args[` + index + `]->value=(void *)(intptr_t)` + name + `;
  // ---- `
	return code
}
