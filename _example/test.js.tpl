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
// }, "aaa")

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
