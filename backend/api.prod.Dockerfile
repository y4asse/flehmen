FROM golang:1.23 as builder
WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
COPY app .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x /app/main
EXPOSE 8080
CMD ["./main"]