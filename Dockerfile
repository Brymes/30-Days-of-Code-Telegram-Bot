FROM golang:1.17.2 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 30DOC .

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /app/30DOC .

ENTRYPOINT ["./30DOC"]