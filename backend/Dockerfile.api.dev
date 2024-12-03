FROM golang:1.23.0

WORKDIR /app

COPY app .

RUN go mod tidy

CMD ["go", "run", "main.go"]