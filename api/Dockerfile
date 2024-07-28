FROM golang:1.23rc2-alpine3.20
RUN apk update && apk add git
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
WORKDIR /api
COPY . .
RUN go mod tidy
RUN go install github.com/air-verse/air@latest
EXPOSE 8888
CMD ["air", "air", "-c", ".air.toml"]