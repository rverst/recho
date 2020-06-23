
FROM golang:1.14-alpine3.12 AS builder

ENV CGO_ENABLED 0
ADD . /go/src/github.com/rverst/request-debug
RUN go build -o /server /go/src/github.com/rverst/request-debug/server.go

FROM alpine:3.12
MAINTAINER "Robert Verst <robert@verst.eu>"

RUN apk add --update --no-cache ca-certificates && \
    mkdir /app

EXPOSE 8080
WORKDIR /app

COPY --from=builder /server ./server

ENV PORT="8080"

CMD ["/app/server"]