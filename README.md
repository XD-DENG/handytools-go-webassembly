# Handy Tools: Written in Golang + WebAssembly + Vue.js

This application is mainly meant to test & demo how to use WebAssembly with Go.

The key steps include:
- Compile Go application (e.g. `wasm/wasm_main.go` in this app) into executable WebAssembly module file, in which you can
    - specify Go functions which can be accessed from JavasScript environment later as JS functions.
    - Operate the objects in JavaScript environment.
    
    Note that you can only compile main packages, otherwise you get an object file that cannot be run in WebAssembly [1].
- Copy the JavaScript support file `GOROOT/misc/wasm/wasm_exec.js` to your project, and load it properly in your HTML page
(ensure the version of Go you use to build is consistent with the version of Go from whose `GOROOT` you copy from).

A live demo can be found at [https://handytools.xd-deng.com](https://handytools.xd-deng.com/)

开发这个应用的主要目的是测试、学习以及展示如何在Go中使用WebAssembly。

关键步骤包括：
- 编译Go程序（例如这里的`wasm/wasm_main.go`）到一个可执行的WebAssembly文件，在其中我们可以：
    - 用Go实现函数，并最终在JavaScript环境中将它们作为JavaScript的函数进行调用
    - 对JavaScript环境中的objects进行操作
    
    注意你只能够编译main packages，否则你只能编译出一个在WebAssembly中不能执行的object文件[1]。
    
- 将Go提供的JavaScript支持文件`GOROOT/misc/wasm/wasm_exec.js`复制到你的项目中，并且在你的HTML文件中正确地调用
（注意你复制这个支持文件的`GOROOT`Go版本应与你用来编译wasm文件的Go版本保持一致） 

Demo请移步[https://handytools.xd-deng.com](https://handytools.xd-deng.com/)

## Commands to Build and Run

```bash
# Compile wasm file
cd wasm
GOARCH=wasm GOOS=js go build -o index.wasm wasm_main.go
cd ..

# Copy the wasm_exec.js file from GOROOT, to double-ensure consistency
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/js/wasm_exec.js

# Start the service
go run main.go
```

## Build & Run with Docker

```bash
docker build -t xddeng/handytools-wasm .
docker run -p 8000:8000 xddeng/handytools-wasm
```

## Reference

- [1] https://github.com/golang/go/wiki/WebAssembly#getting-started
- [2] https://talks.godoc.org/github.com/chai2010/awesome-go-zh/chai2010/chai2010-golang-wasm.slide