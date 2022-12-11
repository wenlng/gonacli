package argsync

func GenDoubleArgTypeCode(name string, index string) string {
	code := `
  double ` + name + ` = 0.0;
  if (wg_info.Length() > ` + index + `) {
    ` + name + ` = wg_info[` + index + `].As<Number>().DoubleValue();
  }`

	return code
}
