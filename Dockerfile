FROM golang:1.17.9-alpine as builder
WORKDIR /go/src/github.com/nguyennh4522/hello-world

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o main main.go

EXPOSE 8080
CMD ["/go/src/github.com/nguyennh4522/hello-world/main"]