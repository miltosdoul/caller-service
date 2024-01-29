FROM golang:alpine3.18 as builder

WORKDIR /opt

ENV COUNTER_SERVICE_URL=$COUNTER_SERVICE_URL \
    LISTEN_PORT=$LISTEN_PORT

COPY . .

RUN go build -o caller-service main.go

FROM alpine:latest

COPY --from=builder /opt/caller-service .

CMD ./caller-service