<div align="center">
<img width="120" style="padding-top: 50px" src="http://47.104.180.148/gonacli/gonacli_logo.svg"/>
<h1 style="margin: 0; padding: 0">GonaCli</h1>
<p>构建生成 NodeJS Addon 扩展的 NAPI C/C++ 中间代码桥接 Golang 的开发工具</p>
<a href="https://goreportcard.com/report/github.com/wenlng/gonacli"><img src="https://goreportcard.com/badge/github.com/wenlng/gonacli"/></a>
<a href="https://goproxy.cn/stats/github.com/wenlng/gonacli/badges/download-count.svg"><img src="https://goproxy.cn/stats/github.com/wenlng/gonacli/badges/download-count.svg"/></a>
<a href="https://godoc.org/github.com/wenlng/gonacli"><img src="https://godoc.org/github.com/wenlng/gonacli?status.svg"/></a>
<a href="https://github.com/wenlng/gonacli/releases"><img src="https://img.shields.io/github/v/release/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli/blob/master/LICENSE"><img src="https://img.shields.io/github/license/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/stars/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/last-commit/wenlng/gonacli.svg"/></a>
</div>

<br/>

> [English](README.md) | 中文

<p>
GONACLI, 发布于 2022 年 12 月。
</p>
<p>
gonacli 开发工具根据配置快速构建生成 NodeJS Addon 扩展的 Napi、C/C++ 中间代码桥接 Golang，在开发 NodeJS Addon 时只需专注在 Golang 代码上的开发。
</p>


<br/>
<p> ⭐️ 如果能帮助到你，记得随手给点一个star。</p>

- Github：[https://github.com/wenlng/gonacli](https://github.com/wenlng/gonacli)

## 中国 Go 代理的列表，任选其一
- ChinaProxy：https://goproxy.cn
- GoProxy https://github.com/goproxy/goproxy.cn
- AliProxy： https://mirrors.aliyun.com/goproxy/
- OfficialProxy： https://goproxy.io/
- Other：https://gocenter.io

## 设置 Go 的代理
- Window
```shell script
$ set GOPROXY=https://goproxy.cn/,direct

### The Golang 1.13+ can be executed directly
$ go env -w GOPROXY=https://goproxy.cn/,direct
```
- Linux or Mac
```shell script
$ export GOPROXY=https://goproxy.cn/,direct

### or
$ echo "export GOPROXY=https://goproxy.cn/,direct" >> ~/.profile
$ source ~/.profile
```

## 使用 golang 方式安装
安装前需要确保系统配置好了 GOPATH 及最终编译到 bin 目录的相关环境变量
``` shell script
# .bash_profile
export GOPATH="/Users/awen/go"
# 配置 bin 目录，使用 golang 方式安装是必须的
export PATH="$PATH:$GOPATH:$GOPATH/bin"
``` 

安装 gonacli 工具
```shell script
$ go install github.com/wenlng/gonacli
$ gonacli --version
```
<br/>

## gonacli 中的命令
### 1、generate 命令
> 根据 goaddon 的配置生成对应 NodeJS Addon 扩展的 Napi、C/C++ 中间代码，用于桥接 Golang 的程序
``` shell script
# 默认将读取当前目录下的 goaddon.json 配置文件
$ gonacli generate

# --config 参数指定配置文件
$ gonacli generate --config demoaddon.json
```
### 2、build 命令
> 相当于 go build -buildmode=c-archive 命令，编译静态库
``` shell script
# 将 Go CGO 编译生成静态库
$ gonacli build

# --args 参数指定 go build 的参数
$ gonacli build --args '-ldflags "-s -w"'
```
### 3、make 命令
> 相当于 node-gyp configure && node-gyp build
命令，将 Napi、C/C++ 代码和静态库编译成最终的 NodeJS Addon 扩展

``` text
使用 make 命令请请确保系统已安装了 node-gyp 编译工具
使用 -npm-i 参数时请确保系统已安装了 NPM 包依赖管理工具
```

``` shell script
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

``` shell script
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
``` shell script
# 保存到 ./demoaddon/ 目录下
$ gonacli build
```

#### 3、生成中间桥接的 Napi C/C++ 代码
``` shell script
# 生成保存到 ./demoaddon/ 目录下
$ gonacli generate --config ./goaddon.json
```

#### 4、编译 Nodejs Adddon
``` shell script
# 生成保存到 ./demoaddon/build 目录下
$ gonacli make --npm-i
```

#### 4、编写 js 测试文件
/test.js
``` javascript
const demoaddon = require('./demoaddon')

const name = "awen"
const res = demoaddon.hello(name)
console.log('>>> ', res)

```

``` shell script
$ node ./test.js
# >>> hello, awen
```

<br/>

## 配置文件
``` json
{
  "name": "demoaddon",      // Nodejs Addon 扩展的名称      
  "sources": [              // go build 的文件列表，注意不能带有路径  
    "demoaddon.go"
  ],
  "output": "./demoaddon/", // 最终输出目录路径
  "exports": [              // 导出的接口，生成 Addon 的 Napi、C/C++ 代码
    {
      "name": "Hello",      // Golang 对应的 //export Hello 接口名称，必须一致
      "args": [             // 传递的参数列表，参数型必须按照下面参照表保持一致
        {                   // 参数要细心严谨，往往是因为配置的类型与 Golang 入口的不一致而导致编译失败
          "name": "name",
          "type": "string"
        }
      ],
      "returntype": "string",   // 返回给 JavaScript 的类型，没有 callback 类型
      "jscallname": "hello",    // JavaScript 调用的名称
      "jscallmode": "sync"      // sync 为同步执行、async 为异步执行（async值必须在args参数中指明 callback 类型参数）
    }
  ]
}
```

## 参数类型对照表

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

### 关于配置文件的 returntype 字段类型
``` text
returntype 字段没有 callback 类型
```

### 关于 array 类型<返回时值有多层时，在 returntype 中不推荐使用>
``` text
1、array 类型在 Golang 接收是字符串类型，需要配合使用 make([]interface{}, 0) 和 json.Unmarshal
2、array 类型在 Golang 返回时是 *C.char 类型，配合使用 json.Marshal
3、array 类型在 JavaScript 传递时是数组类型，但在接收时目前只支持一层，在 Golang 返回多层请使用字符串方式返回再使用 JavaScrpt 的 JSON.parse
```

### 关于 object 类型<返回时值有多层时，在 returntype 中不推荐使用>
``` text
1、object 类型在 Golang 接收是字符串类型，需要配合使用 make([string]interface{}, 0) 和 json.Unmarshal
2、object 类型在 Golang 返回时是 *C.char 类型，配合使用 json.Marshal
3、object 类型在 JavaScript 传递时是数组类型，但在接收时目前只支持一层，在 Golang 返回多层请使用字符串方式返回再使用 JavaScrpt 的 JSON.parse
```

<br/>

> 请作者喝咖啡：[http://witkeycode.com/sponsor](http://witkeycode.com/sponsor)

<br/>

## LICENSE
    MIT
