package main

import "C"
import (
	"encoding/json"
	"fmt"
	"time"
)

//export IntSum32
func IntSum32(x, y int32) C.int {
	// 传入 int32 返回 int32
	return C.int(x + y)
}

//export IntSum64
func IntSum64(x, y int64) C.longlong {
	// 传入 int64  返回 int64
	return C.longlong(x + y)
}

//export UintSum32
func UintSum32(x, y uint32) C.uint {
	// 传入 uint32  返回 uint32
	return C.uint(x + y)
}

//export CompareInt
func CompareInt(x, y int32) bool {
	// 传入 int32  返回 boolean
	return x > y
}

//export FloatSum
func FloatSum(x, y float32) C.float {
	// 传入 float  返回 float
	return C.float(x + y)
}

//export DoubleSum
func DoubleSum(x, y float64) C.double {
	// 传入 double  返回 double
	return C.double(x + y)
}

//export FormatStr
func FormatStr(s *C.char) *C.char {
	// 传入 string  返回 string
	ss := C.GoString(s)
	return C.CString("golang out >>> " + ss)
}

//export EmptyString
func EmptyString(s *C.char) bool {
	// 传入 string  返回 boolean
	ss := C.GoString(s)
	return len(ss) <= 0
}

//export FilterMap
func FilterMap(s *C.char) *C.char {
	// 传入 object  返回 object
	ss := C.GoString(s)
	fmt.Println(fmt.Sprintf("golang out >>> %s", string(ss)))

	var m2 = make(map[string]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>> err: ", err2.Error())
	}
	fmt.Println("golang out >>> map len: ", m2, len(m2))

	var m = make(map[string]string, 2)
	m["a"] = "aaaaa"
	m["b"] = "bbbbb"
	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("golang out >>> err: ", err.Error())
	}

	result := string(jsonStr)
	return C.CString(result)
}

//export CountMap
func CountMap(s *C.char) C.int {
	// 传入 object  返回 int32
	ss := C.GoString(s)
	fmt.Println(fmt.Sprintf("golang out >>> %s", string(ss)))

	var m2 = make(map[string]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>>err: ", err2.Error())
		return 0
	}

	return C.int(len(m2))
}

//export IsMapType
func IsMapType(s *C.char) bool {
	// 传入 object  返回 boolean
	ss := C.GoString(s)
	fmt.Println(fmt.Sprintf("golang out >>> %s", string(ss)))

	var m2 = make(map[string]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>>err: ", err2.Error())
		return false
	}
	return true
}

//export FilterSlice
func FilterSlice(s *C.char) *C.char {
	// 传入 array  返回 array
	ss := C.GoString(s)
	fmt.Println("golang out >>> slice len: ", ss, len(ss))

	var m2 = make([]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>>err: ", err2.Error())
	}
	fmt.Println("golang out >>> slice len: ", m2, len(m2))

	var m = make([]interface{}, 2)
	m[0] = "hello"
	m[1] = "wold"

	jsonStr, err := json.Marshal(m)
	if err != nil {
		fmt.Println("golang out >>>err: ", err.Error())
	}

	result := string(jsonStr)

	return C.CString(result)
}

//export CountSlice
func CountSlice(s *C.char) C.int {
	// 传入 array  返回 int32
	ss := C.GoString(s)
	fmt.Println("golang out >>> slice len: ", ss, len(ss))

	var m2 = make([]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>>err: ", err2.Error())
	}
	fmt.Println("golang out >>> slice len: ", m2, len(m2))

	return C.int(len(m2))
}

//export IsSliceType
func IsSliceType(s *C.char) bool {
	// 传入 object  返回 boolean
	ss := C.GoString(s)
	fmt.Println(fmt.Sprintf("golang out >>> %s", string(ss)))

	var m2 = make([]interface{}, 0)
	err2 := json.Unmarshal([]byte(string(ss)), &m2)
	if err2 != nil {
		fmt.Println("golang out >>>err: ", err2.Error())
		return false
	}
	return true
}

type CallbackOutput struct {
	Data   string `json:"data"`
	Output string `json:"output"`
}

var callbackCount = 0

// ===========================
// 同步执行

//export  SyncCallbackReStr
func SyncCallbackReStr(arg *C.char) *C.char {
	// 同步执行，会发生主线程阻塞
	// 传入 string，返回 string
	result := ""
	ch := make(chan bool)

	go func() {
		callbackCount++
		curCount := callbackCount

		fmt.Println("golang out >>> run", curCount, C.GoString(arg))
		time.Sleep(time.Duration(2) * time.Second)
		var co CallbackOutput
		co.Data = "你好 wait return hello"

		if curCount%2 == 1 {
			co.Data = "你好 wait return wold"
		}

		co.Output = fmt.Sprintf("%d", curCount)
		jsonStr, err := json.Marshal(co)
		if err != nil {
			fmt.Println("golang out >>>err: ", err.Error())
		}
		result = string(jsonStr)
		ch <- true
	}()

	<-ch
	return C.CString(result)
}

//export SyncCallbackReArr
func SyncCallbackReArr(arg *C.char) *C.char {
	// 同步执行，会发生主线程阻塞
	// 传入 array，返回 array
	result := ""
	ch := make(chan bool)
	ss := C.GoString(arg)

	go func() {
		callbackCount++
		curCount := callbackCount

		fmt.Println("golang out >>>> run", curCount, ss)

		var m2 = make([]interface{}, 0)
		err2 := json.Unmarshal([]byte(string(ss)), &m2)
		if err2 != nil {
			fmt.Println("err: ", err2.Error())
		}
		fmt.Println("golang out >>> slice len: ", m2, len(m2))

		var m = make([]interface{}, 3)
		m[0] = "hello"
		m[1] = "wold"
		m[2] = curCount

		jsonStr, err := json.Marshal(m)
		if err != nil {
			fmt.Println("golang out >>>err: ", err.Error())
		}

		result = string(jsonStr)
		ch <- true
	}()

	<-ch
	return C.CString(result)
}

//export SyncCallbackReObject
func SyncCallbackReObject(arg *C.char) *C.char {
	// 同步执行，会发生主线程阻塞
	// 传入 object，返回 object
	result := ""
	ch := make(chan bool)
	ss := C.GoString(arg)

	go func() {
		callbackCount++
		curCount := callbackCount

		fmt.Println("golang out >>> run", curCount, C.GoString(arg))

		var m2 = make(map[string]interface{}, 0)
		err2 := json.Unmarshal([]byte(string(ss)), &m2)
		if err2 != nil {
			fmt.Println("err: ", err2.Error())
		}
		fmt.Println("golang out >>> map len: ", m2, len(m2))

		var m = make(map[string]interface{}, 3)
		m["k1"] = "hello"
		m["k2"] = "wold"
		m["k3"] = curCount

		jsonStr, err := json.Marshal(m)
		if err != nil {
			fmt.Println("golang out >>>err: ", err.Error())
		}

		result = string(jsonStr)
		ch <- true
	}()

	<-ch
	return C.CString(result)
}

//export SyncCallbackReCount
func SyncCallbackReCount(arg *C.char) C.int {
	// 同步执行，会发生主线程阻塞
	// 传入 string，返回 int32
	result := 0
	ch := make(chan bool)

	go func() {
		callbackCount++
		curCount := callbackCount

		fmt.Println("golang out >>> run", curCount, C.GoString(arg))

		result = curCount
		ch <- true
	}()

	<-ch
	return C.int(result)
}

//export SyncCallbackReBool
func SyncCallbackReBool(arg *C.char) bool {
	// 同步执行，会发生主线程阻塞
	// 传入 string，返回 boolean
	result := false
	ch := make(chan bool)

	go func() {
		callbackCount++
		curCount := callbackCount

		fmt.Println("golang out >>> run", curCount, C.GoString(arg))

		result = true
		ch <- true
	}()

	<-ch
	return result
}

//export SyncCallbackSleep
func SyncCallbackSleep(t int32) bool {
	// 同步执行，会发生主线程阻塞
	// 传入 int32，返回 boolean
	ch := make(chan bool)

	go func() {
		callbackCount++
		curCount := callbackCount
		d := t
		time.Sleep(time.Duration(d) * time.Second)
		fmt.Println("golang out >>> run", curCount, d)
		ch <- true
	}()

	<-ch
	return true
}

// =========== 异步

//export  ASyncCallbackReStr
func ASyncCallbackReStr(arg *C.char, cbFuncStr *C.char) *C.char {
	// 异步执行，不会阻塞主线程
	// 传入 string string，返回 string
	return SyncCallbackReStr(arg)
}

//export ASyncCallbackReIntSum64
func ASyncCallbackReIntSum64(x, y int64, cbFuncStr *C.char) C.longlong {
	// 传入 int64  返回 int64
	SyncCallbackSleep(1)
	return C.longlong(x + y)
}

//export ASyncCallbackReUintSum32
func ASyncCallbackReUintSum32(x, y uint32, cbFuncStr *C.char) C.uint {
	// 传入 uint32  返回 uint32
	SyncCallbackSleep(1)
	return C.uint(x + y)
}

//export ASyncCallbackReArr
func ASyncCallbackReArr(arg *C.char, cbFuncStr *C.char) *C.char {
	// 异步执行，不会阻塞主线程
	// 传入 string string，返回 array
	fmt.Println("golang out >>> cbFuncStr = ", C.GoString(cbFuncStr))
	return SyncCallbackReArr(arg)
}

//export ASyncCallbackReObject
func ASyncCallbackReObject(arg *C.char, cbFuncStr *C.char) *C.char {
	// 异步执行，不会阻塞主线程
	// 传入 object string，返回 object
	fmt.Println("golang out >>> cbFuncStr = ", C.GoString(cbFuncStr))
	return SyncCallbackReObject(arg)
}

//export ASyncCallbackReCount
func ASyncCallbackReCount(arg *C.char, cbFuncStr *C.char) C.int {
	// 异步执行，不会阻塞主线程
	// 传入 string string ，返回int32
	fmt.Println("golang out >>> cbFuncStr = ", C.GoString(cbFuncStr))
	return SyncCallbackReCount(arg)
}

//export ASyncCallbackReBool
func ASyncCallbackReBool(arg *C.char, cbFuncStr *C.char) bool {
	// 异步执行，不会阻塞主线程
	// 传入 string string，返回 boolean
	fmt.Println("golang out >>> cbFuncStr = ", C.GoString(cbFuncStr))
	return SyncCallbackReBool(arg)
}

//export ASyncCallbackMArg
func ASyncCallbackMArg(arg *C.char, cbFuncStr *C.char, ext *C.char) bool {
	// 异步执行，不会阻塞主线程
	// 传入 string string，返回 boolean
	fmt.Println("golang out >>> cbFuncStr = ", C.GoString(cbFuncStr))
	return SyncCallbackReBool(arg)
}

func main() {
	// ...
}
