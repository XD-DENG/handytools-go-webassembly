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
- 编译Go程序（例如这里的`wasm/wasm_main.go`），并生成一个可执行的WebAssembly文件，在其中我们可以：
    - 用Go实现函数，并最终在JavaScript环境中将它们作为JavaScript函数进行调用
    - 对JavaScript环境中的objects进行操作
    
    需要注意的是，你只能够编译main packages，否则编译出的object文件将不能在WebAssembly中执行[1]。
    
- 将Go提供的JavaScript支持文件`GOROOT/misc/wasm/wasm_exec.js`复制到你的项目中，同时在你的HTML文件中载入该js文件（这里`GOROOT`的Go版本应与你之前用来编译wasm文件的Go版本保持一致）。

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

## Check in Browser Console

Once the application is running, you can open the corresponding link in your browser
(or directly visit the demo [https://handytools.xd-deng.com](https://handytools.xd-deng.com/))
and open the console as well.

The functions we create in `wasm/wasm_main.go` and expose in the `main()` function are accessible in the browser console,
and we can invoke them in exactly the same way how we invoke JavaScript functions.

<p align="center"> 
    <img src="/static/images/handytools_console.png" alt="drawing" width="80%"/>
</p>


## Reference

- [1] https://github.com/golang/go/wiki/WebAssembly#getting-started
- [2] https://talks.godoc.org/github.com/chai2010/awesome-go-zh/chai2010/chai2010-golang-wasm.slide
