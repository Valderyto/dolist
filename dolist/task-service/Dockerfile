FROM golang:1.20-alpine
WORKDIR /app
COPY main.go ./
RUN go build -o task-service main.go
CMD ["./task-service"]