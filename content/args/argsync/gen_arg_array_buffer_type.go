package argsync

func GenArrayBufferArgTypeCode(name string, index string) (string, string) {
	code := `
  ArrayBuffer wg__` + name + ` = ArrayBuffer();
  if (wg_info.Length() > ` + index + `) {
    wg__` + name + ` = wg_info[` + index + `].As<ArrayBuffer>();
  }
  char * wg_` + name + ` = (char *)wg__` + name + `.Data();
  GoSlice ` + name + ` = wg_build_go_slice(wg_` + name + `, wg__s.ByteLength(), wg__s.ByteLength());`

	//endCode := tools.FormatCodeIndentLn(`delete [] _`+name+`;`, 2)
	//endCode += tools.FormatCodeIndentLn(`delete [] `+name+`;`, 2)
	endCode := ""
	return code, endCode
}
