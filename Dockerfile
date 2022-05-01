FROM --platform=$BUILDPLATFORM alpine:3.15 AS builder

ARG TARGETARCH

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct

WORKDIR $GOPATH/src/QLTools

# 安装必要环境
RUN \
  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache --update go go-bindata g++ ca-certificates tzdata

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .

# 打包项目文件
RUN \
  go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/... && \
  go build -ldflags "-s -w" -o QLTools-linux-$TARGETARCH . && \
  

FROM alpine:3.15

MAINTAINER HomeNavigation "nuanxinqing@gmail.com"

ARG TARGETARCH
ENV TARGETARCH=$TARGETARCH

WORKDIR /QLTools

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder $GOPATH/src/QLTools/QLTools-linux-$TARGETARCH $GOPATH/src/QLTools/config $GOPATH/src/QLTools/docker-entrypoint.sh .

EXPOSE 15000

ENTRYPOINT ["sh", "docker-entrypoint.sh"]
