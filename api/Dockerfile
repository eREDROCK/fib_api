#本番環境用
FROM golang:1.23rc2-alpine3.20

RUN apk update && apk add --no-cache git

ENV TZ=Asia/Tokyo

WORKDIR /api

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8888

CMD ["./main"]
