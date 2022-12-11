package async

// 生成线程执行完成代码
func genWorkCompleteCode(workCompleteName string, structDataName string) string {
	code := `
// -------- genworkThreadCompleteCode
static void ` + workCompleteName + `(napi_env wg_env, napi_status wg_status, void* wg_data) {
  ` + structDataName + `* wg_addon = (` + structDataName + `*)wg_data;
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  wg_catch_err(wg_env, napi_delete_async_work(wg_env, wg_addon->work));
  wg_addon->work = NULL;
  wg_addon->tsfn = NULL;
  for (int i = 0; i < wg_addon->argc; i++) {
    if (wg_addon->args[i]->type == 1) {
      WgAddonArgInfo* info = (WgAddonArgInfo*)wg_addon->args[i];
      delete [] (char *)info->value;
    }
    free(wg_addon->args[i]);
    wg_addon->args[i] = NULL;
  }
}`
	return code
}
