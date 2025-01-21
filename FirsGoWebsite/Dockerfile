FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o myapp .

FROM alpine:latest

WORKDIR /root/

COPY  --from=builder /app/myapp .
COPY --from=builder /app/templates ./templates
EXPOSE 9000

CMD [ "./myapp" ]