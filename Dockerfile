# 第一阶段：构建阶段
FROM golang:1.21.3 as builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .


# 第二阶段：运行阶段
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# 从构建器阶段复制编译后的可执行文件
COPY --from=builder /app/myapp .
COPY --from=builder /app/web .web/
CMD ["./myapp"]
