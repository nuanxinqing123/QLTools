FROM amd64/golang:latest AS builder

MAINTAINER HomeNavigation "nuanxinqing@gmail.com"

WORKDIR $GOPATH/src/QLTools

COPY . .

ADD . ./

# Setting up the AMD64 environment
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# Processing package
RUN go mod tidy

RUN go build -o QLTools-linux-amd64 .

FROM scratch

COPY --from=builder go/src/QLTools/config /config
COPY --from=builder go/src/QLTools/QLTools-linux-amd64 /

ENTRYPOINT ["./QLTools-linux-amd64"]