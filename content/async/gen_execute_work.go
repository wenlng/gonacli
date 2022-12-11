package async

import (
	"node_addon_go/config"
)

// 生成线程执行代码
func genExecuteWorkCode(export config.Export, executeWorkName string, structDataName string) string {
	handlerCode, endCode := genHandlerCode(export)

	code := `
// -------- genExecuteworkCode
static void ` + executeWorkName + `(napi_env wg_env, void* wg_data) {
  ` + structDataName + `* wg_addon = (` + structDataName + `*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));` + handlerCode + `
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));` + endCode + `
}`

	return code
}
