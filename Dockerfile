FROM  golang:1.18.3-alpine AS builder
COPY Tar1 /src
WORKDIR /app
RUN GOPROXY=https://GOPROXY.cn CGO_ENABLE=0 GOOS=linux go build -o api-login /src/main.go
FROM alpine:latest as final
COPY --from=builder /app/api-login /app/
WORKDIR /app/
EXPOSE 9000
CMD ["./api-login"]
# docker build -t first-img .
# docker run -d -p 9000:9000 --name api-login first-img