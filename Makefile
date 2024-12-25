.PHONY: all build run gotool clear help

BINARY="bluebell"
all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/${BINARY}

run:
	@go run ./main.go conf/config.yaml

gotool:
	@go fmt ./
	@go vet ./

clear:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化Go代码, 并编译生成二进制文件"
	@echo "make build - 编译Go代码, 生成二进制文件"
	@echo "make run - 运行Go代码"
	@echo "make gotool - 格式化Go代码"
	@echo "make clear - 清除二进制文件"