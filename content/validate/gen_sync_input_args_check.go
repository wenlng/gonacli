package validate

import "fmt"

func GenRequireWithIndexCode(length int) string {
	ls := fmt.Sprintf("%d", length+1)
	return `
  if(wg_info.Length() < ` + ls + `){
    TypeError::New(wg_env, "The ` + ls + ` parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckNumberWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsNumber()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckBooleanWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsBoolean()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of boolean!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckStringWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsString()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of string!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckArrayWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsArray()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of array!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckArrayBufferWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsArrayBuffer()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of array buffer!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckObjectWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsObject()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of object!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}

func GenCheckFunctionWithIndexCode(index int) string {
	i := fmt.Sprintf("%d", index)
	argPos := fmt.Sprintf("%d", index+1)
	return `
  if(wg_info.Length() > ` + i + ` && !wg_info[` + i + `].IsFunction()){
    TypeError::New(wg_env, "The ` + argPos + ` parameter must be of function!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }`
}
