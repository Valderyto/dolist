FROM golang:1.20-alpine
WORKDIR /app
COPY main.go ./
RUN go build -o user-service main.go
CMD ["./user-service"]
