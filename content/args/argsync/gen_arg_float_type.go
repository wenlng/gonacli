package argsync

func GenFloatArgTypeCode(name string, index string) string {
	code := `
  float ` + name + ` = 0.0;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Number>().FloatValue();
  }`
	return code
}
