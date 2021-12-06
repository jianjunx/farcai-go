FROM golang:alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/bin
WORKDIR /bin

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 main
RUN go build -ldflags="-s -w" -o main ./main.go

###############################################################################
#                                RUN
###############################################################################

FROM alpine:latest

# 添加I18N多语言文件、静态文件、配置文件、模板文件
COPY ./i18n /i18n
COPY ./public /public
COPY ./config /config
COPY ./template /template

COPY --from=builder /bin/main /

###############################################################################
#                                   START
###############################################################################

ENTRYPOINT ["/main"]
