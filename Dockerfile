FROM golang:latest

WORKDIR /app

COPY . /app/

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -gcflags="all=-N -l" -o /app/server /app/cmd/

CMD ["/app/server"]