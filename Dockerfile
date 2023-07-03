FROM golang:1.18-alpine as build

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
COPY .env .

RUN go build -o main main.go

FROM alpine:3.12
RUN apk --update add ca-certificates

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 80

CMD ["/app/main"]