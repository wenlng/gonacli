package validate

import "fmt"

func GenAsyncRequireWithIndexCode(length int) string {
	i := fmt.Sprintf("%d", length)
	argPos := fmt.Sprintf("%d", length+1)
	return `
  if(wg_v` + i + `.IsUndefined()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckNumberWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsNumber()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of number!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckBooleanWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsBoolean()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of boolean!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckStringWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsString()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckArrayWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsArray()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of array!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckArrayBufferWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsArrayBuffer()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of array buffer!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckObjectWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsObject()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of object!").ThrowAsJavaScriptException();
    return NULL;
  }`
}

func GenAsyncCheckFunctionWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(!wg_v` + i + `.IsUndefined() && !wg_v` + i + `.IsFunction()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }`
}
