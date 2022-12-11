const goaddon = require('.')

// console.log(goaddon.intSum32(50, 3))
// console.log(goaddon.intSum64(50, 30))

// console.log(goaddon.compareInt(50, 40))
// console.log(goaddon.compareInt(50, 140))

// console.log(goaddon.floatSum(50, 140.4))
// console.log(goaddon.doubleSum(150, 140.4))

// console.log(goaddon.formatStr("hello"))

// console.log(goaddon.emptyString())
// console.log(goaddon.emptyString("hello"))

const map = {"aa": "bb", "cc": "dd"}
// console.log(goaddon.filterMap())
// console.log(goaddon.countMap(map))
// console.log(goaddon.isMapType(map))

const slice = ["aa", "bb", "cc", "dd"]
// console.log(goaddon.filterSlice(slice))
// console.log(goaddon.countSlice(slice))
// console.log(goaddon.isSliceType(slice))

const content = "sdfddfdf你好1"
const byte = strToArrayBuffer(content).buffer
//
// const resByte = goaddon.filterArrayBuffer(byte)
// console.log(arrayBufferUTF8ToStr(resByte))

// console.log(goaddon.countArrayBuffer(byte))
// console.log(goaddon.isArrayBuffer(byte))

// console.log(goaddon.asyncCallbackSleep(2))

// goaddon.asyncCallbackReStr("awen", function (str) {
//     console.log(">>>>>>", str)
// })

// goaddon.asyncCallbackReUintSum32(100, 200, function (count) {
//     console.log(">>>>>>", count)
// })

// goaddon.asyncCallbackReArr(["hello", "world"], function (arr) {
//     console.log(">>>>>>", arr)
// })

// goaddon.asyncCallbackReArrBuffer(byte, function (resByte) {
//     console.log(">>>>>>", arrayBufferUTF8ToStr(resByte))
// })

const obj = {"name": "awen", "age": "24"}
// goaddon.asyncCallbackReObject(obj, function (resObj) {
//     console.log(">>>>>>", resObj)
// })

// goaddon.asyncCallbackReCount("hello", function (resStr) {
//     console.log(">>>>>>", resStr)
// })

// goaddon.asyncCallbackReBool("hello", function (resBool) {
//     console.log(">>>>>>", resBool)
// })

// goaddon.asyncCallbackMArg("hello", function (resBool) {
//     console.log(">>>>>>", resBool)
// }, byte)

// 对象将会被处理成此种方式传递到 golang map 中
// 但 golang 返回 map 时只处理第一层对象，如果有多层时，建议使用 arraybuffer 或 string 从 golang 返回
// {"name":"Awen_obj","age":"20","list":["ccc","bbb"],"obj":{"aaa":"aaa","ddd":"ddd"},"obj2":{"aaa":"aaa","ddd":"ddd"}}
// const mR = goaddon.filterMap({
//     name: 'Awen_obj',
//     age: 20,
//     list: ["ccc", "bbb"],
//     obj: {
//         aaa: "aaa",
//         ddd: "ddd"
//     },
//     obj2: {
//         aaa: "aaa",
//         ddd: "ddd"
//     }
// })
// console.log(mR)


// 数组将会被处理成此种方式传递到 golang slice 中
// 但 golang 返回 slice 时只处理第一层数组，如果有多层时，建议使用 arraybuffer 或 string 从 golang 返回
//  array ["Awen_arr","20","awen","1997","[object Object]"]
//  array2 ["Awen_arr","20",["awen","1997"],"[object Object]"]
// const sR = goaddon.filterSlice([
//     'Awen_arr',
//     20,
//     [
//         "awen",
//         1997,
//         [
//             "wen",
//             1995
//         ]
//     ],
//     {
//         b: "dd",
//         a: "aaaa"
//     }
// ])
// console.log(sR)


// =========================================================
// 辅助函数
// =========================================================
// String 转 Array Buffer
function strToArrayBuffer(str) {
    const buffer = []
    for (let i of str) {
        const _code = i.charCodeAt(0)
        if (_code < 0x80) {
            buffer.push(_code)
        } else if (_code < 0x800) {
            buffer.push(0xc0 + (_code >> 6))
            buffer.push(0x80 + (_code & 0x3f))
        } else if (_code < 0x10000) {
            buffer.push(0xe0 + (_code >> 12))
            buffer.push(0x80 + (_code >> 6 & 0x3f))
            buffer.push(0x80 + (_code & 0x3f))
        }
    }
    return Uint8Array.from(buffer)
}

// Array Buffer 转 String
function arrayBufferUTF8ToStr(array) {
    let out,i,len,c
    let char2,char3

    if (array instanceof ArrayBuffer) {
        array = new Uint8Array(array)
    }

    out = ""
    len = array.length
    i = 0
    while(i < len) {
        c = array[i++]
        switch(c >> 4) {
            case 0: case 1: case 2: case 3: case 4: case 5: case 6: case 7:
                // 0xxxxxxx
                out += String.fromCharCode(c)
                break;
            case 12: case 13:
                // 110x xxxx   10xx xxxx
                char2 = array[i++];
                out += String.fromCharCode(((c & 0x1F) << 6) | (char2 & 0x3F))
                break
            case 14:
                // 1110 xxxx  10xx xxxx  10xx xxxx
                char2 = array[i++]
                char3 = array[i++]
                out += String.fromCharCode(((c & 0x0F) << 12) |
                    ((char2 & 0x3F) << 6) |
                    ((char3 & 0x3F) << 0))
                break
        }
    }

    return out
}