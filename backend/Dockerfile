# ビルド用イメージ
FROM golang:1.23 as builder

# 作業ディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# ビルドコマンドを実行
RUN go build -o main .

# 実行用イメージ
FROM alpine:3.18

# 必要なディレクトリとバイナリをコピー
WORKDIR /app
COPY --from=builder /app/main .

# ポートを指定
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]