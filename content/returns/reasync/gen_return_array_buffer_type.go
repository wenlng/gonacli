package reasync

import (
	"github.com/wenlng/gonacli/tools"
	"strings"
)

func GenAsyncReturnArrayBufferTypeCode() string {
	return tools.FormatCodeIndentLn(`const void* wg_res_ = (void*)wg_data;`, 2)
}

func GenAsyncCallReturnArrayBufferTypeCode(methodName string, argNames []string) string {
	code := `
  // -------- genHandlerCode
  const void* wg_res_ = ` + methodName + `(` + strings.Join(argNames, ",") + `);`
	return code
}

func GenAsyncCallbackArgArrayBufferTypeCode() string {
	code := `char *wg_ab = (char*) wg_res_;
    size_t wg_ab_length = strlen(wg_ab);
    ArrayBuffer wg_arr_buffer = ArrayBuffer::New(wg_env, wg_ab, wg_ab_length);
    napi_value wg_argv[] = { wg_arr_buffer };`
	return code
}
