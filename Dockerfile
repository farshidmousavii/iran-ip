FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o iran-ip ./cmd/

FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/iran-ip .

RUN adduser -D -u 1000 appuser && \
    mkdir -p /app/data && \
    chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

HEALTHCHECK --interval=1m --timeout=5s --start-period=10s \
  CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/iran-ip"]
CMD ["-addr", ":8080"]
