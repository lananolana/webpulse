FROM golang:1.22.4-alpine AS builder

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
COPY ./internal ./internal

RUN apk update && apk add --no-cache

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/root/.cache/go-build go build -o ./webpulse_backend ./cmd/main.go

FROM alpine:3.20

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

ENV TZ=Europe/Moscow

WORKDIR /app

COPY --from=builder /build/webpulse_backend ./webpulse_backend
COPY ./docs ./docs

VOLUME /app/logs
VOLUME /app/configs

CMD ["./webpulse_backend"]