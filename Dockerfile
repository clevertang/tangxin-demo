FROM golang:latest

# 设置 Go 模块代理
ENV GOPROXY https://goproxy.cn,direct

# 创建工作目录并切换到该目录
WORKDIR /go/src/github.com/tangxin/tangxin-demo

# 将当前目录内容复制到工作目录中
COPY . .

# 编译 Go 程序
RUN go build -o tangxin-demo .

# 暴露端口
EXPOSE 8000

# 指定入口点
ENTRYPOINT ["./tangxin-demo"]
