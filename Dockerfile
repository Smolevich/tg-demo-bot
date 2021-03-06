FROM golang:alpine as build

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

WORKDIR /usr/src/app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . .

RUN go build -o bot-api

FROM alpine

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/bot-api /usr/src/app/bot-api

ENTRYPOINT ["/usr/src/app/bot-api"]