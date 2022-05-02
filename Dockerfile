FROM  alpine:3.15 AS builder

ARG TARGETARCH

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /usr/src/QLTools

# 安装项目必要环境
RUN \
  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache --update go go-bindata g++ ca-certificates tzdata

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . .

# 打包项目文件
RUN \
  #go-bindata -o=bindata/bindata.go -pkg=bindata ./assets/... && \
  go build -ldflags '-linkmode external -s -w -extldflags "-static"' -o QLTools-linux-$TARGETARCH
  

FROM alpine:3.15

MAINTAINER HomeNavigation "nuanxinqing@gmail.com"

ARG TARGETARCH
ENV TARGET_ARCH=$TARGETARCH

WORKDIR /QLTools

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/QLTools/QLTools-linux-$TARGETARCH /usr/src/QLTools/docker-entrypoint.sh /usr/src/QLTools/sample ./

EXPOSE 15000

ENTRYPOINT ["sh", "docker-entrypoint.sh"]