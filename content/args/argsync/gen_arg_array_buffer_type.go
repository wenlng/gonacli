package argsync

func GenArrayBufferArgTypeCode(name string, index string) (string, string) {
	code := `
  ArrayBuffer wg__` + name + ` = ArrayBuffer();
  if (wg_info.Length() > ` + index + `) {
    wg__` + name + ` = wg_info[` + index + `].As<ArrayBuffer>();
  }
  byte * wg_` + name + ` = (byte *)wg__` + name + `.Data();
  char * ` + name + ` = (char *)wg_` + name + `;`

	//endCode := tools.FormatCodeIndentLn(`delete [] _`+name+`;`, 2)
	//endCode += tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	endCode := ""
	return code, endCode
}
