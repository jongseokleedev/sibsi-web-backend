#FROM golang:1.18-alpine as build
#
#WORKDIR /app
#
#ENV GO111MODULE=on
#ENV CGO_ENABLED=0
#ENV GOOS=linux
#ENV GOARCH=amd64
#
#COPY . .
#COPY .env .
#
#RUN go build -o main main.go
#
#FROM alpine:3.12
#RUN apk --update add ca-certificates
#
#WORKDIR /app
#
#COPY --from=build /app/main .
#
#EXPOSE 80
#
#CMD ["/app/main"]
#
#
#FROM golang:alpine AS builder
#
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
#WORKDIR /build
#
#COPY . ./
#
#RUN go mod download
#
#RUN go build -o main .
#
## .env 파일을 복사하여 이미지에 포함
#COPY .env /dist/.env
#
#WORKDIR /dist
#
#
#RUN cp /build/main .
#
#
#FROM scratch
#
#COPY --from=builder /dist/main .
#COPY --from=builder /dist/.env .
#
#ENTRYPOINT ["/main"]

#FROM golang:alpine
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#RUN go mod download
#
#COPY . .
#
## .env 파일을 복사하여 이미지에 포함
#COPY .env .
#
#RUN GOOS=linux GOARCH=amd64 go build -o main .
#
#ENV PORT=80
#ENV GIN_MODE=release
## 다른 환경 변수도 필요한 경우 여기에 추가
#
#ENTRYPOINT ["./main"]


FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .

RUN GOARCH=amd64 GOOS=linux go build -a -o main main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY .env .

COPY --from=builder /app/main .

ENV PORT=8080
ENV GIN_MODE=release
# 다른 환경 변수도 필요한 경우 여기에 추가

CMD ["./main"]

