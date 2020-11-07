FROM golang:1.14 as base

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/app main.go

FROM alpine as final

WORKDIR /usr/local/bin
COPY --from=base /usr/local/bin/app .

EXPOSE 8080

CMD ["app"]