FROM golang:1.15

WORKDIR /go
COPY ./simple-service.go .

RUN go build -o app

EXPOSE 8080/tcp
ENTRYPOINT [ "/go/app" ]