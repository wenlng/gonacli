<div align="center">
<img width="120" style="padding-top: 50px" src="http://47.104.180.148/gonacli/gonacli_logo.svg"/>
<h1 style="margin: 0; padding: 0">GonaCli</h1>
<p>This is a development tool that can quickly use Golang to develop and build NodeJS Addon extensions</p>
<a href="https://goreportcard.com/report/github.com/wenlng/gonacli"><img src="https://goreportcard.com/badge/github.com/wenlng/gonacli"/></a>
<a href="https://godoc.org/github.com/wenlng/gonacli"><img src="https://godoc.org/github.com/wenlng/gonacli?status.svg"/></a>
<a href="https://github.com/wenlng/gonacli/releases"><img src="https://img.shields.io/github/v/release/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli/blob/master/LICENSE"><img src="https://img.shields.io/github/license/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/stars/wenlng/gonacli.svg"/></a>
<a href="https://github.com/wenlng/gonacli"><img src="https://img.shields.io/github/last-commit/wenlng/gonacli.svg"/></a>
</div>

<br/>

> English | [中文](README_zh.md)

<p>
GONACLI is a development tool that quickly uses Golang to develop NodeJS Addon. You only need to concentrate on the development of Golang, and you don't need to care about the implementation of the bridge layer. It supports JavaScript synchronous calls and asynchronous callbacks.
</p>

<br/>

<p> ⭐️ If it helps you, please give a star.</p>

- [https://github.com/wenlng/gonacli](https://github.com/wenlng/gonacli)


<br/>

## Use Golang Install
Ensure that the system is configured with GOPATH environment variables before installation
``` shell
# .bash_profile
export GOPATH="/Users/awen/go"
# set bin dir
export PATH="$PATH:$GOPATH:$GOPATH/bin"
```

Install
``` shell
$ go install github.com/wenlng/gonacli
$ gonacli version
```
<br/>

## Gonacli Command
### 1. generate
Generate Napi, C/C++ bridge code related to NodeJS Addon extension according to the configuration of goaddon
``` shell
# By default, it reads the goaddon in the current directory Json configuration file
$ gonacli generate

# --config Specify Profile
$ gonacli generate --config demoaddon.json
```
### 2. build
Same as the "go build - buildmode=c-archive" command, compile the library
``` shell
# Compile to generate library
$ gonacli build

# --args: Specify the args of go build
$ gonacli build --args '-ldflags "-s -w"'
```
### 3. make
Same as the "node-gyp configure && node-gyp build" command，Compile NodeJS Addon

``` text
Please ensure that the node gyp compiler has been installed on the system before using the "make" command

Before using the "--npm-i" arg, ensure that the system has installed the NPM package dependency management tool
```

``` shell
# --npm-i: Use NPM Instal Napi and Bindings dependency
# --npm-i: Before run npm install，Then run node-gyp configure && node-gyp build
$ gonacli make --npm-i

# --npm-i: It can be directly executed "make" after using npm to install dependencies
$ gonacli make

# --args: Specify the parameters of node gyp build，for example "--debug"
$ gonacli make --args '--debug'
```

<br/>

## Quick Use
<p>Tip：Ensure that relevant commands can be used normally</p>

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


#### 1. Create Goaddon Configure File
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

#### 2. Write Golang Code
/demoaddon.go
``` go
package main
import "C"

// notice：//export xxxx is necessary

//export Hello
func Hello(_name *C.char) s *C.char {
	// args string type，return string type
	name := C.GoString(_name)
	
	res := "hello"
	if len(name) > 0 {
	    res += "," + name
	}
	
	return C.CString(res)
}
```

Compile libraries
``` shell
# Save to the "./demoaddon/" directory
$ gonacli build
```

#### 3. Generate bridging Napi C/C++code
``` shell
# Save to the "./demoaddon/" directory
$ gonacli generate --config ./goaddon.json
```

#### 4. Compile Nodejs Adddon
``` shell
# Save to the "./demoaddon/build" directory
$ gonacli make --npm-i
```

#### 5. Create JS Test File
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

## Configure File Description
``` text
{
  "name": "demoaddon",      // Name of Nodejs Addon
  "sources": [              // File list of go build，Cannot have path
    "demoaddon.go"
  ],
  "output": "./demoaddon/", // Output directory path
  "exports": [              // Exported interface, generating the Napi and C/C++ code of Addon
    {
      "name": "Hello",      // The name of the "//export Hello" interface corresponding to Golang must be consistent
      "args": [             // The parameter type of the passed parameter list must be consistent with the type table
        {                  
          "name": "name",
          "type": "string"
        }
      ],
      "returntype": "string",   // The type returned to JavaScript，has no callback type
      "jscallname": "hello",    // JavaScript call name
      "jscallmode": "sync"      // Sync is synchronous execution, and Async is asynchronous execution
    }
  ]
}
```

## Type Table

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

### The returntype field type of the configuration file
<p>The returntype field has no callback type</p>

### array type
<p>When there are multiple levels when returning, it is not recommended to use in the returntype</p>
<p>1. The "array" type received in Golang is a string "*C.Char" type, which needs to be use "make([]interface{}, 0)" and "json.Unmarshal"</p>
<p>2. The "array" type is when Golang returns "*C.Char" type, use "json.Marshal"</p>
<p>3. The "array" type is an Array type when JavaScript is passed, but currently only supports one layer when receiving. Please use string method to return multiple layers in Golang, and then use JavaScript's "JSON.parse"</p>

### object type
<p>When there are multiple levels when returning, it is not recommended to use in the returntype</p>
<p>1. The "object" type received in Golang is a string type. You need to use "make([string]interface{}, 0)" and "json.Unmarshal"</p>
<p>2. 2. The "object" type is when Golang returns "*C.Char" type, use "json.Marshal"</p>
<p>3. The "object" type is an Object type when JavaScript is passed, but currently only supports one layer when receiving. Please use string method to return multiple layers in Golang, and then use JavaScript's "JSON.parse"</p>

<br/>

## JavaScript Sync Call
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

#### 2. Golang Code
/demoaddon.go
``` go
package main
import "C"

//export Hello
func Hello(_name *C.char) s *C.char {
	// args is string type，return string type
	name := C.GoString(_name)
	
	res := "hello"
	ch := make(chan bool)

	go func() {
	    // Time consuming task processing
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

## JavaScript Async Call
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

#### 2. Golang Code
/demoaddon.go
``` go
package main
import "C"

//export Hello
func Hello(_name *C.char, cbsFnName *C.char) s *C.char {
	// args is string type，return string type
	name := C.GoString(_name)
	
	res := "hello"
	ch := make(chan bool)

    go func() {
	    // Time consuming task processing
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

## LICENSE
    MIT
