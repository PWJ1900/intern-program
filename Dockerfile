FROM  golang:1.18.3-alpine AS builder
COPY Tar3 /src
WORKDIR /app
RUN GOPROXY=https://GOPROXY.cn CGO_ENABLE=0 GOOS=linux go mod init Tar3 &&  go get github.com/go-ldap/ldap && go build -o api-login-ldap /src/ldapUse.go /src/requestOrg.go
FROM alpine:latest as final
COPY --from=builder /app/api-login-ldap /app/
WORKDIR /app/
EXPOSE 9000
CMD ["./api-login-ldap"]
# docker build -t first-img .
# docker run -d -p 9000:9000 --name api-login first-img