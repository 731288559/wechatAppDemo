#源镜像
FROM golang:latest

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

#设置工作目录
WORKDIR /Users/chenjiayu/Documents/code/practice/wechatAppDemo

#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/github.com/wechatAppDemo

#go构建可执行文件
RUN go build -o myDemo gin_demo.go

#暴露端口
EXPOSE 8080

#最终运行docker的命令
ENTRYPOINT  ["./myDemo"]