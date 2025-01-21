FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o myapp .

FROM alpine:latest

WORKDIR /root/

COPY  --from=builder /app/myapp .

EXPOSE 9000

CMD [ "./myapp" ]