package base

import (
	"node_addon_go/config"
	"node_addon_go/tools"
)

func genWgAddonDataCode() string {
	return `
//---------- genWgAddonArg ----------
typedef struct {
  int type; // [1]char [2]int [3]float [4]double [5]bool [6]byte
  int len;
  void* value;
} WgAddonArgInfo;`
}

func genBuildGoStringCode() string {
	return `
//---------- genBuildGoString ----------
GoString _buildGoString(const char* p, size_t n){
  return {p, static_cast<ptrdiff_t>(n)};
}`
}

func genCatchErrCode() string {
	return `
// ------------- genCatchErr -----------
static void wg_catch_err(napi_env env, napi_status status) {
  if (status != napi_ok) {
    const napi_extended_error_info* error_info = NULL;
    napi_get_last_error_info(env, &error_info);
    printf("addon >>>>> %s\n", error_info->error_message);
    exit(0);
  }
}`
}

func genStringSplitCode() string {
	return `
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
}`
}

func genArrayToStringCode() string {
	return `
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
}`
}

func genStringToArrayCode() string {
	return `
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
}`
}

func genObjectArrToStringCode() string {
	return `
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
}`
}

func genObjectToStringCode() string {
	return `
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
}`
}

func genStringToObject() string {
	code := `
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
}`

	return code
}

func GenBeforeCode(hasAsync bool) string {
	code := `// [common]++++++++++++++++++++++++++++++++++++++ start`
	code += genWgAddonDataCode()
	//code += genBuildGoStringCode()
	code += genStringSplitCode()
	code += genArrayToStringCode()
	code += genStringToArrayCode()

	code += genObjectArrToStringCode()
	code += genObjectToStringCode()
	code += genStringToObject()

	if hasAsync {
		code += genCatchErrCode()
	}

	code += `
// [common]++++++++++++++++++++++++++++++++++++++ end`
	return code
}

func genExportsJsCallApi(exports []config.Export) string {
	code := ""
	for _, export := range exports {
		jsApiName := export.JsCallName
		goApiName := export.Name
		if export.JsCallMode == "sync" {
			code += tools.FormatCodeIndentLn(`exports.Set(String::New(env, "`+jsApiName+`"), Function::New(env, _`+goApiName+`));`, 2)
		}
	}

	return code
}

func GenAfterCode(cfg config.Config, asyncCode string) string {
	name := cfg.Name

	exportsCode := genExportsJsCallApi(cfg.Exports)
	code := `
Object Init(Env env, Object exports) {` + exportsCode + asyncCode + `
  return exports;
}

NODE_API_MODULE(` + name + `, Init)`
	return code
}
