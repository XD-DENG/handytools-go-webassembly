FROM golang:1.13.5-alpine3.10 AS builder

WORKDIR /app
COPY . /app
ENV GOPATH=/app

# Build wasm file
RUN cd wasm;GOARCH=wasm GOOS=js go build -o index.wasm wasm_main.go

# Copy the wasm_exec.js file from the GOROOT of the image, to further ensure consistency
RUN cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static/js/wasm_exec.js
RUN go build -o bin/main main.go

# For smaller image size
# see https://medium.com/@gdiener/how-to-build-a-smaller-docker-image-76779e18d48a
FROM alpine:3.10

WORKDIR /app

COPY --from=builder /app/static ./static
COPY --from=builder /app/wasm ./wasm
COPY --from=builder /app/bin ./bin

CMD ["./bin/main"]
