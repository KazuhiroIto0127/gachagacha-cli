FROM golang:1.21.1-alpine

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

WORKDIR /app
COPY ./ /app

RUN apk add --no-cache build-base make

RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/spf13/cobra-cli@latest
