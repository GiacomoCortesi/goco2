FROM golang:1.21.2-alpine3.17 AS builder
LABEL MAINTAINER="Giacomo Cortesi <giacomo.cortesi1993@gmail.com>"

WORKDIR /app/src

# Copy sources
COPY cmd/ cmd/
COPY *.go .
COPY go.mod .
COPY go.sum .

RUN CGO_ENABLED=0 go build -o /app/build/goco2 /app/src/cmd/main.go

FROM alpine:3.14
WORKDIR /app

COPY static/ static/
ENV HTTP_HOST 0.0.0.0
ENV HTTP_PORT 8000

COPY --from=builder /app/build/goco2 .

CMD ["./goco2"]
