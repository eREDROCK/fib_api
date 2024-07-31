#本番環境用
# ベースイメージの指定
FROM golang:1.23rc2-alpine3.20

# 必要なパッケージのインストール
RUN apk update && apk add --no-cache git

# タイムゾーンの設定
ENV TZ=Asia/Tokyo

# 作業ディレクトリの設定
WORKDIR /api

# ソースコードのコピー
COPY . .

# モジュールの依存関係を解決
RUN go mod tidy

# バイナリのビルド
RUN go build -o main .

# ポートの公開
EXPOSE 8888

# コンテナ起動時のコマンド
CMD ["./main"]
