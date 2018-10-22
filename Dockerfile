FROM golang:1.10 as builder
WORKDIR $GOPATH/src/github.com/johan-lejdung/go-microservice-api-template
COPY ./ .
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp go-microservice-api-template /go-service

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /go-service /
ENTRYPOINT [ "/go-service" ]
