<div align="center">
<img width="120" style="padding-top: 50px" src="http://47.104.180.148/gonacli/gonacli_logo.svg"/>
<h1 style="margin: 0; padding: 0">GonaCli</h1>
<p>一个快速使用 Golang 编写和构建生成 NodeJS Addon 扩展的开发工具</p>
<a href="https://goreportcard.com/report/github.com/wenlng/gonacli"><img src="https://goreportcard.com/badge/github.com/wenlng/gonacli"/></a>
<a href="https://godoc.org/github.com/wenlng/gonacli"><img src="https://godoc.org/github.com/wenlng/gonacli?status.svg"/></a>
<a href="https://github.com/wenlng/gonacli/releases"><img src="https://img.shields.io/github/v/release/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli/blob/master/LICENSE"><img src="https://img.shields.io/github/license/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/stars/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/last-commit/wenlng/gonacli.svg"/></a>
</div>

<br/>

> [English](README.md) | 中文

<br/>

<p>
GONACLI 是一个快速使用 Golang 开发 NodeJS Addon 扩展的开发工具，开发者只需要专注在 Golang 的开发，无需关心与 NodeJS 的 Bridge 桥接层的实现，支持 JavaScript 同步调用和异步回调等。
</p>


<br/>
<p> ⭐️ 如果能帮助到你，记得随手给点一个star。</p>

- [https://github.com/wenlng/gonacli](https://github.com/wenlng/gonacli)

- Gonacli QQ群(技术支持)：885267905

## Gonacli 的兼容支持
- Linux
- Mac OS
- Windows

## NodeJS Addon 的兼容支持
- Linux / Mac OS / Windows
- NodeJS(12.0+)
- Npm(6.0+)
- Node-gyp(9.0+)
- Go(1.14+)

## 使用 go 方式安装 gonacli 工具
<p>安装前需要确保系统配置好了 GOPATH 及最终编译保存到 bin 目录的环境变量</p>

Linux or Mac OS
``` shell
# .bash_profile
export GOPATH="/Users/awen/go"

# 配置 bin 目录，使用 golang 方式安装是必须的
export PATH="$PATH:$GOPATH:$GOPATH/bin"
``` 

Windows
``` shell
# 打开系统环境变量设置
GOPATH: C:\awen\go
# 配置 bin 目录，使用 golang 方式安装是必须的
PATH: %GOPATH%\bin
``` 

开始安装

```shell
# linux or Mac OS
$ GOPROXY=https://goproxy.cn/,direct && go install github.com/wenlng/gonacli@latest

# widow
$ set GOPROXY=https://goproxy.cn/,direct && go install github.com/wenlng/gonacli@latest

$ gonacli version
```
<br/>

## Windows 环境编译
<p> 在 Windows 开发环境下需要安装 Go CGO 需要的 gcc/g++ 编译器，可以下载 "MinGW" 安装，配置好 MinGW/bin 的 PATH 环境变量即可，在命令行能够正常执行 gcc 。</p>

``` shell
$ gcc -v
```

<p>Window 环境下还需要安装 NodeJS Addon 编译工具 node-gyp 依赖的 c/c++ 编译工具</p>

``` shell
$ npm install --global --production windows-build-tools
```

<br/>

## GONACLI 中的命令参数

### 1、generate

<p>根据 goaddon 的配置生成对应 NodeJS Addon 相关的 Napi、C/C++ 桥接代码</p>

``` shell
# 默认将读取当前目录下的 goaddon.json 配置文件
$ gonacli generate

# --config 参数指定配置文件
$ gonacli generate --config demoaddon.json
```

### 2、build

<p>相当于 go build -buildmode=c-archive 命令，编译静态库</p>

``` shell
# 将 Go CGO 编译生成静态库
$ gonacli build

# --config 参数指定配置文件
# --args 参数指定 go build 的参数，需要用 '' 引号包裹
$ gonacli build --args '-ldflags "-s -w"'
```

### 3、install

<p>相当于 npm install 命令， 安装 NodeJS 需要的相关依赖</p>

``` shell
# --config 参数指定配置文件
$ gonacli install --config demoaddon.json
```

### 4、msvc

<p>该命令只针对 window 环境下的兼容处理，需要 dlltool.exe 或 lib.exe (二选一)</p>
<p>1、"MinGW" 支持 "dlltool.exe" 工具</p>
<p>2、"Microsoft Visual c++ Build tools" 或 "Visual Studio" 的 "lib.exe" 工具</p>

``` shell
# --vs 参数表示使用 VS 的 "lib.exe" 工具，默认是 MinGW 的 "dlltool.exe" 工具
# --32x 参数表示支持 32 位的系统，默认 64 位
# --config 参数指定配置文件
$ gonacli msvc --config demoaddon.json
```

### 5、make
<p>相当于 node-gyp configure && node-gyp build 命令，编译成最终的 NodeJS Addon 扩展</p>

<p>使用 make 命令请请确保系统已安装了 node-gyp 编译工具</p>

``` shell
# 编译
$ gonacli make

# --args 参数指定 node-gyp build 的参数，例如调试 --debug 参数
$ gonacli make --args '--debug'
```


<br/>

## 使用 Golang 快速开发 NodeJS Addon 的 Demo

<p>Tip：请确保相关命令能正常使用，该 Demo 是在 Linux / Macos 环境下进行</p>

``` shell
# go
$ go version

# node
$ node -v

# npm
$ npm -v

# node-gyp
$ node-gyp -v

# gcc
$ gcc -v
```


#### 1、新建配置文件
/goaddon.json
``` json
{
  "name": "demoaddon",
  "sources": [
    "demoaddon.go"
  ],
  "output": "./demoaddon/",
  "exports": [
    {
      "name": "Hello",
      "args": [
        {
          "name": "name",
          "type": "string"
        }
      ],
      "returntype": "string",
      "jscallname": "hello",
      "jscallmode": "sync"
    }
  ]
}
```

#### 2、编写 Golang 代码
/demoaddon.go
``` go
package main
import "C"

// 注意：//export xxxx 是必须的

//export Hello
func Hello(_name *C.char) *C.char {
	// 传入 string 类型，返回 string 类型
	name := C.GoString(_name)
	
	res := "hello"
	if len(name) > 0 {
	    res += "," + name
	}
	
	return C.CString(res)
}
```

#### 3、生成桥接的 Napi C/C++ 代码
``` shell
# 生成保存到 ./demoaddon/ 目录下
$ gonacli generate --config ./goaddon.json
```

#### 4、编译静态库
``` shell
# 保存到 ./demoaddon/ 目录下
$ gonacli build
```

#### 5、安装 Nodejs 相关依赖
``` shell
# 保存到 ./demoaddon/ 目录下
$ gonacli install
```

#### 6、编译 Nodejs Adddon
``` shell
# 生成保存到 ./demoaddon/build 目录下
$ gonacli make
```

#### 7、编写 js 测试文件
/demoaddon/test.js
``` javascript
const demoaddon = require('.')

const name = "awen"
const res = demoaddon.hello(name)
console.log('>>> ', res)

```

``` shell
$ node ./test.js
# >>> hello, awen
```

<br/>

## 配置文件说明
``` text
{
  "name": "demoaddon",      // Nodejs Addon 扩展的名称      
  "sources": [              // go build 的文件列表，注意不能带有路径  
    "demoaddon.go"
  ],
  "output": "./demoaddon/", // 最终输出目录路径
  "exports": [              // 导出的接口，生成 Addon 的 Napi、C/C++ 代码
    {
      "name": "Hello",      // 对应 Golang 的 "//export Hello" 接口名称，必须一致
      "args": [             // 传递的参数列表，参数型必须按照下面的类型表保持一致
        {                   // 参数要细心严谨，往往是因为配置的类型与 Golang 入口的不一致而导致编译失败
          "name": "name",   // 参数名称，但不能与当前参数列表中某一项重复
          "type": "string"  // 参数类型
        }
      ],
      "returntype": "string",   // 返回给 JavaScript 的类型，没有 callback 类型
      "jscallname": "hello",    // JavaScript 调用的名称
      "jscallmode": "sync"      // sync 为同步执行、async 为异步执行（async值必须在args参数中指明 callback 类型参数）
    }
  ]
}
```

## 类型对照表
<p> -------- 请严格按照类型对照表 ------- </p>

|    Type     | Golang Args | Golang Return  |   JS / TS   |
|:-----------:|:-----------:|:--------------:|:-----------:|
|     int     |    int32    |     C.int      |   number    |
|    int32    |    int32    |     C.int      |   number    |
|    int64    |    int64    |   C.longlong   |   number    |
|   uint32    |   uint32    |     C.uint     |   number    |
|    float    |   float32   |    C.float     |   number    |
|   double    |   float64   |    C.double    |   number    |
|   boolean   |    bool     |      bool      |   boolean   |
|   string    |   *C.char   |    *C.char     |   string    |
|    array    |   *C.char   |    *C.char     |    Array    |
|   object    |   *C.char   |    *C.char     |   Object    |
|  callback   |   *C.char   |       -        |  Function   |

### 配置文件的 returntype 字段类型
<p>returntype 字段没有 callback 类型</p>

### array 类型（当返回时存在多层时，在 returntype 中不推荐使用）
<p>1、array 类型在 Golang 接收是字符串类型，需要配合使用 []interface{} 和 json.Unmarshal</p>
<p>2、array 类型在 Golang 返回时是 *C.char 类型，配合使用 json.Marshal</p>
<p>3、array 类型在 JavaScript 传递时是数组类型，但在接收时目前只支持一层，在 Golang 返回多层请使用字符串方式返回再使用 JavaScrpt 的 JSON.parse</p>

### object 类型（当返回时存在多层时，在 returntype 中不推荐使用）
<p>1、object 类型在 Golang 接收是字符串 *C.char 类型，需要配合使用 [string]interface{} 和 json.Unmarshal</p>
<p>2、object 类型在 Golang 返回时是 *C.char 类型，配合使用 json.Marshal</p>
<p>3、object 类型在 JavaScript 传递时是对象类型，但在接收时目前只支持一层，在 Golang 返回多层请使用字符串方式返回再使用 JavaScrpt 的 JSON.parse</p>

<br/>

## JavaScript 同步式调用
/goaddon.json
``` json
{
  "name": "demoaddon",
  "sources": [
    "demoaddon.go"
  ],
  "output": "./demoaddon/",
  "exports": [
    {
      "name": "Hello",
      "args": [
        {
          "name": "name",
          "type": "string"
        }
      ],
      "returntype": "string",
      "jscallname": "hello",
      "jscallmode": "sync"
    }
  ]
}
```

#### 2、编写 Golang 代码
/demoaddon.go
``` go
package main
import "C"

//export Hello
func Hello(_name *C.char) *C.char {
	// 传入 string 类型，返回 string 类型
	name := C.GoString(_name)
	
	res := "hello"
	ch := make(chan bool)

    // 当使用协程时，由于 JS 使用同步式调用，JS 进程会发生阻塞等待返回
	go func() {
	    // 耗时任务处理
	    time.Sleep(time.Duration(2) * time.Second)
		if len(name) > 0 {
	        res += "," + name
	    }   
		ch <- true
	}()

	<-ch
	
	return C.CString(res)
}
```

#### 3. Test
/test.js
``` javascript
const demoaddon = require('./demoaddon')

const name = "awen"
const res = demoaddon.hello(name)
console.log('>>> ', res)
```

<br/>

## JavaScript 异步式回调
/goaddon.json
``` json
{
  "name": "demoaddon",
  "sources": [
    "demoaddon.go"
  ],
  "output": "./demoaddon/",
  "exports": [
    {
      "name": "Hello",
      "args": [
        {
          "name": "name",
          "type": "string"
        },
        {
          "name": "cbs",
          "type": "callback"
        }
      ],
      "returntype": "string",
      "jscallname": "hello",
      "jscallmode": "async"
    }
  ]
}
```

#### 2、编写 Golang 代码
/demoaddon.go
``` go
package main
import "C"

//export Hello
func Hello(_name *C.char, cbsFnName *C.char) *C.char {
	// 传入 string 类型，返回 string 类型
	name := C.GoString(_name)
	
	res := "hello"
	ch := make(chan bool)

    // 当使用协程时，由于 JS 使用异步式调用，JS 进程不会发生阻塞，当返回值时会 JS callback
	go func() {
	    // 耗时任务处理
		time.Sleep(time.Duration(2) * time.Second)
		if len(name) > 0 {
	        res += "," + name
	    }   
		ch <- true
	}()

	<-ch
	
	return C.CString(res)
}
```

#### 3. Test
/test.js
``` javascript
const demoaddon = require('./demoaddon')

const name = "awen"
demoaddon.hello(name, funciton(res){
    console.log('>>> ', res)
})
```

<br/>

> 请作者喝咖啡：[http://witkeycode.com/sponsor](http://witkeycode.com/sponsor)

<br/>

## LICENSE
    MIT
