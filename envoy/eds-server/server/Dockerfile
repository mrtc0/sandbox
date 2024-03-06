FROM golang:1.22.1 as builder

WORKDIR /go/src/app
ENV GO111MODULE=on

RUN groupadd -g 10001 user \
    && useradd -u 10001 -g user user

COPY go.mod go.sum ./

RUN go mod download

COPY eds-sever.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/eds-server

# FROM scratch
# 
# COPY --from=builder /go/bin/eds-server /go/bin/eds-server
# COPY --from=builder /etc/group /etc/group
# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# 
# EXPOSE 20000
# 
# USER user
# 
# ENTRYPOINT ["/go/bin/eds-server"]
ENTRYPOINT ["tail", "-f", "/dev/null"]
