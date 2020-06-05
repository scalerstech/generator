FROM golang:alpine

RUN mkdir /go/src/generator

WORKDIR /go/src/generator

COPY . ./

RUN go env

RUN go build -o ./cmd/generator/generator ./cmd/generator/generator.go

CMD ./cmd/generator/generator