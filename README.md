# Handy Tools: A Toy Service Written in Golang + WebAssembly + Vue.js

This application is mainly meant to test & demo how to use WebAssembly with Go.

The key steps include:
- Compile Go application (e.g. `wasm/wasm_main.go` in this app) into executable WebAssembly module file, in which you can
    - specify Go functions which can be accessed from JavasScript environment later as JS functions.
    - Operate the objects in JavaScript environment.
- Copy the JavaScript support file `GOROOT/misc/wasm/wasm_exec.js` to your project, and load it properly in your HTML page
(ensure the version of Go you use to build is consistent with the version of Go from whose `GOROOT` you copy from).

A live demo can be found at [https://handytools.xd-deng.com](https://handytools.xd-deng.com/)

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

- https://github.com/golang/go/wiki/WebAssembly#getting-started
- https://talks.godoc.org/github.com/chai2010/awesome-go-zh/chai2010/chai2010-golang-wasm.slide