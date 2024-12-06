FROM golang:1.23.0

WORKDIR /app

COPY app .

RUN go mod tidy

RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]