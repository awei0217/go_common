# 基于golang1.12（母镜像） 构建docker镜像
FROM golang:1.12
# 执行一个命令
RUN mkdir -p /go/src/docker_go_demo
# 工作目录
WORKDIR /go/src/docker_go_demo

# 将文件复制到镜像中 . 代表当前目录
ADD . /go/src/docker_go_demo

# 执行操作，跟在终端执行语句一样此处是编译
RUN go build .

# 程序入口
ENTRYPOINT ["./docker_go_demo"]