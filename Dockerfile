FROM golang:1.23-alpine3.20

RUN apk add --no-cache musl-dev gcc make g++ file
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p /app
COPY . /app
WORKDIR /app

ENV GO111MODULE=on
RUN go mod tidy
RUN go get github.com/githubnemo/CompileDaemon
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  ./scripts/seed.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/http/main.go
EXPOSE 8080
CMD ["./app"]