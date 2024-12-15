FROM golang:1.12.0-alpine3.9
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o main .
CMD ["/app/main"]


