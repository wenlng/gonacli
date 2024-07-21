#include <napi.h>
#include <string>
#include <assert.h>
#include <functional>

#include "goaddon.h"

using namespace Napi;
using namespace std;

#define NAPI_EXPERIMENTAL
// [common]++++++++++++++++++++++++++++++++++++++ start
//---------- genWgAddonArg ----------
typedef struct {
  int type; // [1]char [2]int [3]float [4]double [5]bool
  int len;
  void* value;
} WgAddonArgInfo;
// ------------- genStringSplit -----------
void wg_string_split(const string& str, const char split, vector<string>& res){
  if (str == "") return;
  string strs = str + split;
  size_t pos = strs.find(split);
  while (pos != strs.npos){
	string temp = strs.substr(0, pos);
	res.push_back(temp);
	strs = strs.substr(pos + 1, strs.size());
	pos = strs.find(split);
  }
}
// ------------- genStringToArray2 -----------
string wg_array_to_string(Array arr) {
  string res = "[";
  string last;
  for(uint32_t i = 0; i < arr.Length(); i++){
    Value v = arr[i];
    if (last.size() > 0) {
      res += last;
      last = "";
    }
    if (v.IsArray()){
      Array arr2 = v.As<Array>();
      res += wg_array_to_string(arr2);
      last = ",";
    } else {
      string ss = v.ToString();
      res += "\"" + ss + "\"";
      last = ",";
    }
  }
  res += "]";
  return res;
}
// ------------- genStringToArray -----------
Array wg_string_to_array(string str, Env env) {
  Array arr = Array::New(env);
  vector<string> strList;
  if (str == "") return arr;
  size_t pos;
  while ((pos = str.find("[")) != string::npos) {
	str.replace(pos, 1, ",");
  }
  while ((pos = str.find("]")) != string::npos) {
    str.replace(pos, 1, ",");
  }
  wg_string_split(str, ',', strList);
  int index = 0;
  for (auto s : strList) {
    if (s.size() > 0) {
      int _spos = s.find("\"");
      s = s.substr(_spos + 1);
      int _epos = s.find("\"");
      s = s.substr(0, _epos);
      arr.Set(Number::New(env, index), String::New(env, s));
      index++;
    }
  }
  return arr;
}
// ------------- genObjectArrToString -----------
string wg_object_to_string(Object objs);
string wg_object_array_to_string(Array arr) {
  string res = "[";
  string last;
  for(uint32_t i = 0; i < arr.Length(); i++){
    Value v = arr[i];
    if (last.size() > 0) {
      res += last;
      last = "";
    }
    if (v.IsArray()){
      Array arr2 = v.As<Array>();
      res += wg_object_array_to_string(arr2);
      last = ",";
    } else if (v.IsObject()){
      Object obj2 = v.As<Object>();
      res += wg_object_to_string(obj2);
      last = ",";
    } else {
      string ss = v.ToString();
      res += "\"" + ss + "\"";
      last = ",";
    }
  }
  res += "]";
  return res;
}
// ------------- genObjectToString -----------
string wg_object_to_string(Object objs) {
  string res = "{";
  Array keyArr = objs.GetPropertyNames();
  string last;
  for(uint32_t i = 0; i < keyArr.Length(); i++){
    Value key = keyArr[i];
    Value v = objs.Get(key);
    string name = key.As<String>().Utf8Value();
    if (last.size() > 0) {
      res += last;
      last = "";
    }
    res += "\"" + name + "\":";
    if (v.IsArray()) {
      Array arr = v.As<Array>();
      res += wg_object_array_to_string(arr);
      last = ",";
    } else if (v.IsObject()){
      Object obj2 = v.As<Object>();
      res += wg_object_to_string(obj2);
      last = ",";
    } else {
      string ss = v.ToString();
      res += "\"" + ss + "\"";
      last = ",";
    }
  }
  res += "}";
  return res;
}
// ------------- genStringToObject -----------
Object wg_string_to_object(string str, Env env) {
  Object obj = Object::New(env);
  vector<string> strList;
  if (str == "") return obj;
  size_t pos;
  while ((pos = str.find("{")) != string::npos) {
    str.replace(pos, 1, ",");
  }
  while ((pos = str.find("}")) != string::npos) {
    str.replace(pos, 1, ",");
  }
  wg_string_split(str, ',', strList);
  for (auto s : strList) {
    if (s.size() > 0) {
      while ((pos = s.find("\"")) != string::npos) {
        s.replace(pos, 1, "");
      }
      vector<string> keyValue;
      wg_string_split(s, ':', keyValue);
      string key = keyValue.size() >= 0 ? keyValue[0] : "";
      string value = keyValue.size() >= 1 ? keyValue[1] : "";
      obj.Set(String::New(env, key), String::New(env, value));
    }
  }
  return obj;
}
// ------------- genCatchErr -----------
static void wg_catch_err(napi_env env, napi_status status) {
  if (status != napi_ok) {
    const napi_extended_error_info* error_info = NULL;
    napi_get_last_error_info(env, &error_info);
    printf("addon >>>>> %s\n", error_info->error_message);
    exit(0);
  }
}
// [common]++++++++++++++++++++++++++++++++++++++ end
// ---------- GenCode ---------- 
Value _IntSum32(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 1){
    TypeError::New(wg_env, "The 1 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  int x = 0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().Int32Value();
  }
  int y = 0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().Int32Value();
  }
  int wg_res_ = IntSum32(x,y);
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _IntSum64(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 2){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  long long int x = 0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().Int64Value();
  }
  long long int y = 0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().Int64Value();
  }
  long long int wg_res_ = IntSum64(x,y);
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _UintSum32(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 2){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  int x = 0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().Int32Value();
  }
  int y = 0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().Int32Value();
  }
  unsigned long wg_res_ = UintSum32(x,y);
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _CompareInt(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 2){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  int x = 0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().Int32Value();
  }
  int y = 0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().Int32Value();
  }
  bool wg_res_ = CompareInt(x,y);
  return Boolean::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _FloatSum(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 2){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  float x = 0.0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().FloatValue();
  }
  float y = 0.0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().FloatValue();
  }
  float wg_res_ = FloatSum(x,y);
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _DoubleSum(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() < 1){
    TypeError::New(wg_env, "The 1 parameter must be passed!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  if(wg_info.Length() > 1 && !wg_info[1].IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  double x = 0.0;
  if (wg_info.Length() > 0) {
    x = wg_info[0].As<Number>().DoubleValue();
  }
  double y = 0.0;
  if (wg_info.Length() > 1) {
    y = wg_info[1].As<Number>().DoubleValue();
  }
  double wg_res_ = DoubleSum(x,y);
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _FormatStr(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  string wg_s = "";
  if (wg_info.Length() > 0) {
    wg_s = wg_info[0].As<String>().Utf8Value();
  }
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  string wg_res_ = FormatStr(s);
  delete [] s;
  return String::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _EmptyString(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  string wg_s = "";
  if (wg_info.Length() > 0) {
    wg_s = wg_info[0].As<String>().Utf8Value();
  }
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  bool wg_res_ = EmptyString(s);
  delete [] s;
  return Boolean::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _FilterMap(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsObject()){
    TypeError::New(wg_env, "The 1 parameter must be of object!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Object wg__obj = Object::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__obj = wg_info[0].As<Object>();
  }
  string wg_obj = wg_object_to_string(wg__obj);
  char *obj = new char[wg_obj.length() + 1];
  strcpy(obj, wg_obj.c_str());
  string wg_res_ = FilterMap(obj);
  Object wg_obj_ = wg_string_to_object(wg_res_, wg_env);
  delete [] obj;
  return wg_obj_;
}
// ---------- GenCode ---------- 
Value _CountMap(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsObject()){
    TypeError::New(wg_env, "The 1 parameter must be of object!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Object wg__obj = Object::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__obj = wg_info[0].As<Object>();
  }
  string wg_obj = wg_object_to_string(wg__obj);
  char *obj = new char[wg_obj.length() + 1];
  strcpy(obj, wg_obj.c_str());
  int wg_res_ = CountMap(obj);
  delete [] obj;
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _IsMapType(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsObject()){
    TypeError::New(wg_env, "The 1 parameter must be of object!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Object wg__obj = Object::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__obj = wg_info[0].As<Object>();
  }
  string wg_obj = wg_object_to_string(wg__obj);
  char *obj = new char[wg_obj.length() + 1];
  strcpy(obj, wg_obj.c_str());
  bool wg_res_ = IsMapType(obj);
  delete [] obj;
  return Boolean::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _FilterSlice(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsArray()){
    TypeError::New(wg_env, "The 1 parameter must be of array!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Array wg__s = Array::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__s = wg_info[0].As<Array>();
  }
  string wg_s = wg_array_to_string(wg__s);
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  string wg_res_ = FilterSlice(s);
  Array wg_arr = wg_string_to_array(wg_res_, wg_env);
  delete [] s;
  return wg_arr;
}
// ---------- GenCode ---------- 
Value _CountSlice(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsArray()){
    TypeError::New(wg_env, "The 1 parameter must be of array!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Array wg__s = Array::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__s = wg_info[0].As<Array>();
  }
  string wg_s = wg_array_to_string(wg__s);
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  int wg_res_ = CountSlice(s);
  delete [] s;
  return Number::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _IsSliceType(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsArray()){
    TypeError::New(wg_env, "The 1 parameter must be of array!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  Array wg__s = Array::New(wg_env);
  if (wg_info.Length() > 0) {
    wg__s = wg_info[0].As<Array>();
  }
  string wg_s = wg_array_to_string(wg__s);
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  bool wg_res_ = IsSliceType(s);
  delete [] s;
  return Boolean::New(wg_env, wg_res_);
}
// ---------- GenCode ---------- 
Value _SyncCallbackSleep(const CallbackInfo& wg_info) {
  Env wg_env = wg_info.Env();
  if(wg_info.Length() > 0 && !wg_info[0].IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return wg_env.Null();
  }
  int s = 0;
  if (wg_info.Length() > 0) {
    s = wg_info[0].As<Number>().Int32Value();
  }
  bool wg_res_ = SyncCallbackSleep(s);
  return Boolean::New(wg_env, wg_res_);
}
// [ASyncCallbackReStr] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[2];
} WgAddonDataASyncCallbackReStr;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackrestr(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const string wg__res_ = (char*)wg_data;
  if (wg_env != NULL) {
    napi_value wg_res_ = String::New(wg_env, wg__res_);
    napi_value wg_argv[] = { wg_res_ };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackrestr(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReStr* wg_addon = (WgAddonDataASyncCallbackReStr*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  // arg - s
  WgAddonArgInfo * wg_sInfo = wg_addon->args[0];
  char * s = new char[wg_sInfo->len];
  strcpy(s, (char *)wg_sInfo->value);
  // ---- 
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const void* wg_res_ = ASyncCallbackReStr(s,cb);
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] s;
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackrestr(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReStr* wg_addon = (WgAddonDataASyncCallbackReStr*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackrestr(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 2;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[2];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReStr* wg_addon = (WgAddonDataASyncCallbackReStr*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - s
  string wg_s = "";
  if(!wg_v0.IsUndefined() && wg_v0.IsString()){
    wg_s = wg_v0.As<String>().Utf8Value();
  }
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_s.length() + 1;
  wg_addon->args[0]->value=(void *)s;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackrestr, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackrestr, wg_work_complete_asynccallbackrestr, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackrestr(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_str",NULL,wg_work_asynccallbackrestr,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReStr]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackReUintSum32] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[3];
} WgAddonDataASyncCallbackReUintSum32;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackreuintsum32(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const int wg_res_ = (int)(intptr_t)wg_data;
  if (wg_env != NULL) {
    Number wg_int_ = Number::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_int_ };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackreuintsum32(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReUintSum32* wg_addon = (WgAddonDataASyncCallbackReUintSum32*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  // arg - x
  WgAddonArgInfo * wg_xInfo = wg_addon->args[0];
  int x = (int)(intptr_t)wg_xInfo->value;
  // ---- 
  // arg - y
  WgAddonArgInfo * wg_yInfo = wg_addon->args[1];
  int y = (int)(intptr_t)wg_yInfo->value;
  // ---- 
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[2];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const int wg_res__ = ASyncCallbackReUintSum32(x,y,cb);
  void * wg_res_ = (void *)(intptr_t)wg_res__;
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackreuintsum32(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReUintSum32* wg_addon = (WgAddonDataASyncCallbackReUintSum32*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackreuintsum32(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 3;
  size_t wg_cb_arg_index = 2;
  napi_value wg_args[3];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReUintSum32* wg_addon = (WgAddonDataASyncCallbackReUintSum32*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsNumber()){
    TypeError::New(wg_env, "The 1 parameter must be of number!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - x
  int x = 0;
  if(!wg_v0.IsUndefined() && wg_v0.IsNumber()){
    x = wg_v0.As<Number>().Int32Value();
  }
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=2;
  wg_addon->args[0]->value=(void *)(intptr_t)x;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(!wg_v1.IsUndefined() && !wg_v1.IsNumber()){
    TypeError::New(wg_env, "The 2 parameter must be of number!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - y
  int y = 0;
  if(!wg_v1.IsUndefined() && wg_v1.IsNumber()){
    y = wg_v1.As<Number>().Int32Value();
  }
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=2;
  wg_addon->args[1]->value=(void *)(intptr_t)y;
  // ---- 
  Value wg_v2 = Value(wg_env, wg_args[2]);
  if(wg_v2.IsUndefined()){
    TypeError::New(wg_env, "The 3 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v2.IsUndefined() && !wg_v2.IsFunction()){
    TypeError::New(wg_env, "The 3 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v2.IsUndefined() && wg_v2.IsFunction()){
    wg__cb = wg_v2.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[2] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[2]));
  wg_addon->args[2]->type=1;
  wg_addon->args[2]->len=wg__cb.length() + 1;
  wg_addon->args[2]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackreuintsum32, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackreuintsum32, wg_work_complete_asynccallbackreuintsum32, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackreuintsum32(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_uint_sum32",NULL,wg_work_asynccallbackreuintsum32,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReUintSum32]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackReArr] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[2];
} WgAddonDataASyncCallbackReArr;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackrearr(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const string wg_res_ = (char*)wg_data;
  if (wg_env != NULL) {
    Array wg_arr = wg_string_to_array(wg_res_, wg_env);
    napi_value wg_argv[] = { wg_arr };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackrearr(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReArr* wg_addon = (WgAddonDataASyncCallbackReArr*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  WgAddonArgInfo * wg_sInfo = wg_addon->args[0];
  char *s = new char[wg_sInfo->len];
  strcpy(s, (char *)wg_sInfo->value);
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const char* wg_res_ = ASyncCallbackReArr(s,cb);
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] s;
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackrearr(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReArr* wg_addon = (WgAddonDataASyncCallbackReArr*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackrearr(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 2;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[2];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReArr* wg_addon = (WgAddonDataASyncCallbackReArr*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsArray()){
    TypeError::New(wg_env, "The 1 parameter must be of array!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - s
  Array wg__s = Array::New(wg_env);
  if(!wg_v0.IsUndefined() && wg_v0.IsArray()){
    wg__s = wg_v0.As<Array>();
  }
  string wg_s = wg_array_to_string(wg__s);
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_s.length() + 1;
  wg_addon->args[0]->value=(void *)s;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackrearr, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackrearr, wg_work_complete_asynccallbackrearr, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackrearr(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_arr",NULL,wg_work_asynccallbackrearr,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReArr]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackReObject] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[2];
} WgAddonDataASyncCallbackReObject;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackreobject(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const string wg_res_ = (char*)wg_data;
  if (wg_env != NULL) {
    Object wg_obj = wg_string_to_object(wg_res_, wg_env);
    napi_value wg_argv[] = { wg_obj };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackreobject(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReObject* wg_addon = (WgAddonDataASyncCallbackReObject*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  WgAddonArgInfo * wg_sInfo = wg_addon->args[0];
  char *s = new char[wg_sInfo->len];
  strcpy(s, (char *)wg_sInfo->value);
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const char* wg_res_ = ASyncCallbackReObject(s,cb);
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));  delete [] s;
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackreobject(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReObject* wg_addon = (WgAddonDataASyncCallbackReObject*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackreobject(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 2;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[2];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReObject* wg_addon = (WgAddonDataASyncCallbackReObject*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsObject()){
    TypeError::New(wg_env, "The 1 parameter must be of object!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - s
  Object wg__s = Object::New(wg_env);
  if(!wg_v0.IsUndefined() && wg_v0.IsObject()){
    wg__s = wg_v0.As<Object>();
  }
  string wg_s = wg_object_to_string(wg__s);
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_s.length() + 1;
  wg_addon->args[0]->value=(void *)s;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackreobject, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackreobject, wg_work_complete_asynccallbackreobject, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackreobject(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_object",NULL,wg_work_asynccallbackreobject,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReObject]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackReCount] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[2];
} WgAddonDataASyncCallbackReCount;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackrecount(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const int wg_res_ = (int)(intptr_t)wg_data;
  if (wg_env != NULL) {
    Number wg_int_ = Number::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_int_ };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackrecount(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReCount* wg_addon = (WgAddonDataASyncCallbackReCount*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  // arg - s
  WgAddonArgInfo * wg_sInfo = wg_addon->args[0];
  char * s = new char[wg_sInfo->len];
  strcpy(s, (char *)wg_sInfo->value);
  // ---- 
  // arg - cb995
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const int wg_res__ = ASyncCallbackReCount(s,cb);
  void * wg_res_ = (void *)(intptr_t)wg_res__;
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] s;
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackrecount(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReCount* wg_addon = (WgAddonDataASyncCallbackReCount*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackrecount(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 2;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[2];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReCount* wg_addon = (WgAddonDataASyncCallbackReCount*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - s
  string wg_s = "";
  if(!wg_v0.IsUndefined() && wg_v0.IsString()){
    wg_s = wg_v0.As<String>().Utf8Value();
  }
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_s.length() + 1;
  wg_addon->args[0]->value=(void *)s;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackrecount, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackrecount, wg_work_complete_asynccallbackrecount, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackrecount(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_count",NULL,wg_work_asynccallbackrecount,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReCount]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackReBool] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[2];
} WgAddonDataASyncCallbackReBool;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackrebool(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const bool wg_res_ = (bool)wg_data;
  if (wg_env != NULL) {
    Boolean wg_bool_ = Boolean::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_bool_ };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackrebool(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackReBool* wg_addon = (WgAddonDataASyncCallbackReBool*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  // arg - s
  WgAddonArgInfo * wg_sInfo = wg_addon->args[0];
  char * s = new char[wg_sInfo->len];
  strcpy(s, (char *)wg_sInfo->value);
  // ---- 
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const bool wg_res_ = ASyncCallbackReBool(s,cb);
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] s;
  delete [] cb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackrebool(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackReBool* wg_addon = (WgAddonDataASyncCallbackReBool*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackrebool(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 2;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[2];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackReBool* wg_addon = (WgAddonDataASyncCallbackReBool*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(!wg_v0.IsUndefined() && !wg_v0.IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - s
  string wg_s = "";
  if(!wg_v0.IsUndefined() && wg_v0.IsString()){
    wg_s = wg_v0.As<String>().Utf8Value();
  }
  char *s = new char[wg_s.length() + 1];
  strcpy(s, wg_s.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_s.length() + 1;
  wg_addon->args[0]->value=(void *)s;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackrebool, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackrebool, wg_work_complete_asynccallbackrebool, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackrebool(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_re_bool",NULL,wg_work_asynccallbackrebool,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackReBool]+++++++++++++++++++++++++++++++++ end
// [ASyncCallbackMArg] +++++++++++++++++++++++++++++++++ start
typedef struct{
  napi_async_work work;
  napi_threadsafe_function tsfn;
  int argc;
  WgAddonArgInfo *args[3];
} WgAddonDataASyncCallbackMArg;
// ------------ genJsCallbackCode
static void wg_js_callback_asynccallbackmarg(napi_env wg_env, napi_value wg_js_cb, void* wg_context, void* wg_data) {
  (void)wg_context;
  const bool wg_res_ = (bool)wg_data;
  if (wg_env != NULL) {
    Boolean wg_bool_ = Boolean::New(wg_env, wg_res_);
    napi_value wg_argv[] = { wg_bool_ };
    napi_value wg_global;
	napi_get_global(wg_env, &wg_global);
    wg_catch_err(wg_env, napi_call_function(wg_env, wg_global, wg_js_cb, 1, wg_argv, NULL));
  }
}
// -------- genExecuteworkCode
static void wg_execute_workasynccallbackmarg(napi_env wg_env, void* wg_data) {
  WgAddonDataASyncCallbackMArg* wg_addon = (WgAddonDataASyncCallbackMArg*)wg_data;
  wg_catch_err(wg_env, napi_acquire_threadsafe_function(wg_addon->tsfn));
  // arg - str
  WgAddonArgInfo * wg_strInfo = wg_addon->args[0];
  char * str = new char[wg_strInfo->len];
  strcpy(str, (char *)wg_strInfo->value);
  // ---- 
  // arg - cb
  WgAddonArgInfo * wg_cbInfo = wg_addon->args[1];
  char * cb = new char[wg_cbInfo->len];
  strcpy(cb, (char *)wg_cbInfo->value);
  // ---- 
  // arg - arrb
  WgAddonArgInfo * wg_arrbInfo = wg_addon->args[2];
  char * arrb = new char[wg_arrbInfo->len];
  strcpy(arrb, (char *)wg_arrbInfo->value);
  // ---- 
  // -------- genHandlerCode
  const bool wg_res_ = ASyncCallbackMArg(str,cb,arrb);
  wg_catch_err(wg_env, napi_call_threadsafe_function(wg_addon->tsfn, (void*)(wg_res_), napi_tsfn_blocking));
  wg_catch_err(wg_env, napi_release_threadsafe_function(wg_addon->tsfn, napi_tsfn_release));
  delete [] str;
  delete [] cb;
  delete [] arrb;
}
// -------- genworkThreadCompleteCode
static void wg_work_complete_asynccallbackmarg(napi_env wg_env, napi_status wg_status, void* wg_data) {
  WgAddonDataASyncCallbackMArg* wg_addon = (WgAddonDataASyncCallbackMArg*)wg_data;
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
}
// ---------- genworkThreadCode
static napi_value wg_work_asynccallbackmarg(napi_env wg_env, napi_callback_info wg_info) {
  size_t wg_argc = 3;
  size_t wg_cb_arg_index = 1;
  napi_value wg_args[3];
  napi_value wg_work_name;
  napi_status wg_sts;
  WgAddonDataASyncCallbackMArg* wg_addon = (WgAddonDataASyncCallbackMArg*)malloc(sizeof(*wg_addon));
  wg_addon->work = NULL;
  wg_addon->argc = wg_argc;
  wg_sts = napi_get_cb_info(wg_env, wg_info, &wg_argc, wg_args, NULL, NULL);
  wg_catch_err(wg_env, wg_sts);
  napi_value wg_js_cb = wg_args[wg_cb_arg_index];
  Value wg_v0 = Value(wg_env, wg_args[0]);
  if(wg_v0.IsUndefined()){
    TypeError::New(wg_env, "The 1 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v0.IsUndefined() && !wg_v0.IsString()){
    TypeError::New(wg_env, "The 1 parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - str
  string wg_str = "";
  if(!wg_v0.IsUndefined() && wg_v0.IsString()){
    wg_str = wg_v0.As<String>().Utf8Value();
  }
  char *str = new char[wg_str.length() + 1];
  strcpy(str, wg_str.c_str());
  wg_addon->args[0] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[0]));
  wg_addon->args[0]->type=1;
  wg_addon->args[0]->len=wg_str.length() + 1;
  wg_addon->args[0]->value=(void *)str;
  // ---- 
  Value wg_v1 = Value(wg_env, wg_args[1]);
  if(wg_v1.IsUndefined()){
    TypeError::New(wg_env, "The 2 parameter must be passed!").ThrowAsJavaScriptException();
    return NULL;
  }
  if(!wg_v1.IsUndefined() && !wg_v1.IsFunction()){
    TypeError::New(wg_env, "The 2 parameter must be of function!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - cb
  string wg__cb = "cb(){}";
  if(!wg_v1.IsUndefined() && wg_v1.IsFunction()){
    wg__cb = wg_v1.ToString();
    size_t pos = wg__cb.find("{");
    if (pos > 0) {
      wg__cb = wg__cb.substr(0, pos);
      wg__cb += "{}";
    }
  }
  char * cb = new char[wg__cb.length() + 1];
  strcpy(cb, wg__cb.c_str());
  wg_addon->args[1] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[1]));
  wg_addon->args[1]->type=1;
  wg_addon->args[1]->len=wg__cb.length() + 1;
  wg_addon->args[1]->value=(void *)cb;
  // ---- 
  Value wg_v2 = Value(wg_env, wg_args[2]);
  if(!wg_v2.IsUndefined() && !wg_v2.IsString()){
    TypeError::New(wg_env, "The 3 parameter must be of string!").ThrowAsJavaScriptException();
    return NULL;
  }
  // arg - arrb
  string wg_arrb = "";
  if(!wg_v2.IsUndefined() && wg_v2.IsString()){
    wg_arrb = wg_v2.As<String>().Utf8Value();
  }
  char *arrb = new char[wg_arrb.length() + 1];
  strcpy(arrb, wg_arrb.c_str());
  wg_addon->args[2] = (WgAddonArgInfo*)malloc(sizeof(*wg_addon->args[2]));
  wg_addon->args[2]->type=1;
  wg_addon->args[2]->len=wg_arrb.length() + 1;
  wg_addon->args[2]->value=(void *)arrb;
  // ---- 
  assert(wg_addon->work == NULL && "Only one work item must exist at a time");
  wg_catch_err(wg_env, napi_create_string_utf8(wg_env, "N-API Thread-safe Call from Async Work Item", NAPI_AUTO_LENGTH, &wg_work_name));
  wg_sts = napi_create_threadsafe_function(wg_env, wg_js_cb, NULL, wg_work_name, 0, 1, NULL, NULL, NULL, wg_js_callback_asynccallbackmarg, &(wg_addon->tsfn));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_create_async_work(wg_env, NULL, wg_work_name, wg_execute_workasynccallbackmarg, wg_work_complete_asynccallbackmarg, wg_addon, &(wg_addon->work));
  wg_catch_err(wg_env, wg_sts);
  wg_sts = napi_queue_async_work(wg_env, wg_addon->work);
  wg_catch_err(wg_env, wg_sts);
  return NULL;
}
// ---------- GenRegisterAsyncCode ---------- 
void register_asynccallbackmarg(Env env, Object exports){
  napi_property_descriptor desc = {"async_callback_m_arg",NULL,wg_work_asynccallbackmarg,NULL,NULL,NULL,napi_default,NULL};
  napi_define_properties(env, exports, 1, &desc);
}
// [ASyncCallbackMArg]+++++++++++++++++++++++++++++++++ end
Object Init(Env env, Object exports) {
  exports.Set(String::New(env, "int_sum32"), Function::New(env, _IntSum32));
  exports.Set(String::New(env, "int_sum64"), Function::New(env, _IntSum64));
  exports.Set(String::New(env, "uint_sum32"), Function::New(env, _UintSum32));
  exports.Set(String::New(env, "compare_int"), Function::New(env, _CompareInt));
  exports.Set(String::New(env, "float_sum"), Function::New(env, _FloatSum));
  exports.Set(String::New(env, "double_sum"), Function::New(env, _DoubleSum));
  exports.Set(String::New(env, "format_str"), Function::New(env, _FormatStr));
  exports.Set(String::New(env, "empty_string"), Function::New(env, _EmptyString));
  exports.Set(String::New(env, "filter_map"), Function::New(env, _FilterMap));
  exports.Set(String::New(env, "count_map"), Function::New(env, _CountMap));
  exports.Set(String::New(env, "is_map_type"), Function::New(env, _IsMapType));
  exports.Set(String::New(env, "filter_slice"), Function::New(env, _FilterSlice));
  exports.Set(String::New(env, "count_slice"), Function::New(env, _CountSlice));
  exports.Set(String::New(env, "is_slice_type"), Function::New(env, _IsSliceType));
  exports.Set(String::New(env, "async_callback_sleep"), Function::New(env, _SyncCallbackSleep));
  register_asynccallbackrestr(env, exports);
  register_asynccallbackreuintsum32(env, exports);
  register_asynccallbackrearr(env, exports);
  register_asynccallbackreobject(env, exports);
  register_asynccallbackrecount(env, exports);
  register_asynccallbackrebool(env, exports);
  register_asynccallbackmarg(env, exports);
  return exports;
}

NODE_API_MODULE(goaddon, Init)