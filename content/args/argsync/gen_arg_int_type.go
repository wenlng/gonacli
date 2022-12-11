package argsync

func GenIntArgTypeCode(name string, index string, varType string) string {
	// 	int64 uint32 int32
	//  int64_t uint32_t int32_t
	code := ""
	if varType == "int64" {
		code = `
  long long int ` + name + ` = 0;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Number>().Int64Value();
  }`
	} else if varType == "uint32" {
		code = `
  unsigned long ` + name + ` = 0;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Number>().Uint32Value();
  }`
	} else {
		code = `
  int ` + name + ` = 0;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Number>().Int32Value();
  }`
	}

	return code
}
