# ── 阶段一：编译 ──────────────────────────────────
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o resume-api .

# ── 阶段二：运行（极小镜像）─────────────────────────
FROM alpine:3.19

WORKDIR /app

# 复制二进制和数据文件
COPY --from=builder /app/resume-api .
COPY --from=builder /app/data ./data

EXPOSE 8080

CMD ["./resume-api"]
