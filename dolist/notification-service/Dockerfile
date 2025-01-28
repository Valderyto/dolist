FROM golang:1.20-alpine
WORKDIR /app
COPY main.go ./
RUN go build -o notification-service main.go
CMD ["./notification-service"]