npm config set registry https://registry.npmjs.com

npm login -d

git tag -d v1.0.7
git push origin :v1.0.7
git push origin --delete tag v1.0.4

git tag -a v1.0.7 -m "Release version 1.0.7"
git push origin v1.0.7


GOOS=windows go build gonacli.go


// 用gcc从该def和静态库生成一个dll
gcc goaddon.def goaddon.a -shared -lwinmm -lWs2_32 -o goaddon.dll -Wl,--out-implib,goaddon.dll.a

// 需要使用 MSVC 相关工具，请安装如下其中之一
1、Microsoft Visual c++ Build tools
2、Visual Studio 2022

// 用MSVC自带的 lib 程序生成MSVC可用的 .lib 文件：
// /MACHINE:X86, 如果是64位环境，则将最后一个参数改为 /MACHINE:X64 

lib /def:goaddon.def /name:goaddon.dll /out:goaddon.lib /MACHINE:X64

修改Go编译器生成的 libgo.h 文件，将其中2行 line #1 开头的行删掉，2行 _Complex 相关的代码删掉。

这时，MSVC程序可以直接 #include "libgo.h" ，代码中可以直接调用函数 Function1 ，并在链接时添加对 godll.lib 的引用。

// 尝试实验
dlltool -d goaddon.def -D goaddon.dll -l goaddon.lib