<div align="center">
<img width="120" style="padding-top: 50px" src="http://47.104.180.148/gonacli/gonacli_logo.svg"/>
<h1 style="margin: 0; padding: 0">GonaCli</h1>
<p>一套快速使用 Golang 开发和构建生成 NodeJS Addon 扩展的开发工具</p>
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
GONACLI 是一套快速使用 Golang 开发 NodeJS Addon 扩展的开发工具，你只需要专注在 Golang 上的开发，无需关心 Bridge 桥接层的实现，支持 JavaScript 同步调用和异步回调等。
</p>


<br/>
<p> ⭐️ 如果能帮助到你，记得随手给点一个star。</p>

- [https://github.com/wenlng/gonacli](https://github.com/wenlng/gonacli)


## 使用 go 方式安装 gonacli 工具
安装前需要确保系统配置好了 GOPATH 及最终编译保存到 bin 目录的相关环境变量
``` shell
# .bash_profile
export GOPATH="/Users/awen/go"
# 配置 bin 目录，使用 golang 方式安装是必须的
export PATH="$PATH:$GOPATH:$GOPATH/bin"
``` 

安装
```shell
$ GOPROXY=https://goproxy.cn/,direct && go install github.com/wenlng/gonacli
$ gonacli version
```
<br/>

## gonacli 中的命令
### 1、generate
根据 goaddon 的配置生成对应 NodeJS Addon 扩展相关的 Napi、C/C++ 桥接代码
``` shell
# 默认将读取当前目录下的 goaddon.json 配置文件
$ gonacli generate

# --config 参数指定配置文件
$ gonacli generate --config demoaddon.json
```
### 2、build
相当于 go build -buildmode=c-archive 命令，编译静态库
``` shell
# 将 Go CGO 编译生成静态库
$ gonacli build

# --args 参数指定 go build 的参数
$ gonacli build --args '-ldflags "-s -w"'
```
### 3、make
相当于 node-gyp configure && node-gyp build 命令，编译成最终的 NodeJS Addon 扩展

``` text
使用 make 命令请请确保系统已安装了 node-gyp 编译工具
使用 -npm-i 参数时请确保系统已安装了 NPM 包依赖管理工具
```

``` shell
# --npm-i 参数是使用 NPM 安装 Napi 和 Bindings 依赖
# --npm-i 参数等同于先执行 npm install，后再执行 node-gyp configure && node-gyp build
$ gonacli make --npm-i

# --npm-i 参数在首次执行 make 时指定即可，第二次 make 后因为安装过依赖无需再次指定
# 直接执行 node-gyp configure && node-gyp build 编译扩展
$ gonacli make

# --args 参数指定 node-gyp build 的参数，例如调试 --debug 参数
$ gonacli make --args '--debug'
```

<br/>

## 快速使用
<p>Tip：确保相关命令能正常使用</p>

``` shell
# go
$ go version

# node
$ node -v

# npm
$ npm -v

# node-gyp
$ node-gyp -v
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
import "C"

// 注意：//export xxxx 是必须的

//export Hello
func Hello(_name *C.char) s *C.char {
	// 传入 string 类型，返回 string 类型
	name := C.GoString(_name)
	
	res := "hello"
	if len(name) > 0 {
	    res += "," + name
	}
	
	return C.CString(res)
}
```

编译静态库
``` shell
# 保存到 ./demoaddon/ 目录下
$ gonacli build
```

#### 3、生成桥接的 Napi C/C++ 代码
``` shell
# 生成保存到 ./demoaddon/ 目录下
$ gonacli generate --config ./goaddon.json
```

#### 4、编译 Nodejs Adddon
``` shell
# 生成保存到 ./demoaddon/build 目录下
$ gonacli make --npm-i
```

#### 5、编写 js 测试文件
/test.js
``` javascript
const demoaddon = require('./demoaddon')

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
| arraybuffer |   *C.char   | unsafe.Pointer | ArrayBuffer |
|  callback   |   *C.char   |       -        |  Function   |

### 配置文件的 returntype 字段类型
<p>returntype 字段没有 callback 类型</p>

### array 类型（当返回时存在多层时，在 returntype 中不推荐使用）
<p>1、array 类型在 Golang 接收是字符串类型，需要配合使用 make([]interface{}, 0) 和 json.Unmarshal</p>
<p>2、array 类型在 Golang 返回时是 *C.char 类型，配合使用 json.Marshal</p>
<p>3、array 类型在 JavaScript 传递时是数组类型，但在接收时目前只支持一层，在 Golang 返回多层请使用字符串方式返回再使用 JavaScrpt 的 JSON.parse</p>

### object 类型（当返回时存在多层时，在 returntype 中不推荐使用）
<p>1、object 类型在 Golang 接收是字符串 *C.char 类型，需要配合使用 make([string]interface{}, 0) 和 json.Unmarshal</p>
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
import "C"

//export Hello
func Hello(_name *C.char) s *C.char {
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
import "C"

//export Hello
func Hello(_name *C.char, cbsFnName *C.char) s *C.char {
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

<br/>

> 请作者喝咖啡：[http://witkeycode.com/sponsor](http://witkeycode.com/sponsor)

<br/>

## LICENSE
    MIT
