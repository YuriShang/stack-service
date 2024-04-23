FROM golang:1.22.2-alpine3.18 AS builder

WORKDIR /usr/local/go/src/

ADD stackService/ /usr/local/go/src/

RUN go clean --modcache
RUN go build -mod=readonly -o ./ ./main.go

FROM alpine:3.18

COPY --from=builder /usr/local/go/src/ /
COPY --from=builder /usr/local/go/src/config.yml /

CMD ["/main"]