package argsync

func GenBooleanArgTypeCode(name string, index string) string {
	code := `
  bool ` + name + ` = false;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Boolean>();
  }`
	return code
}
