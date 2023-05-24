FROM golang:1.20.3-buster

RUN go version
ENV GOPATH=/

COPY ./ ./


RUN go mod download
RUN go build -o todo ./cmd/main.go

CMD ["./todo"]