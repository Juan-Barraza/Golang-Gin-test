FROM golang:1.26.4-alpine AS builder

WORKDIR /app

COPY go.mod  go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-execute ./cmd/api/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates && \
    update-ca-certificates

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/go-execute /app/go-execute

RUN chown appuser:appgroup /app/go-execute && \
    chmod 550 /app/go-execute

USER appuser

EXPOSE 8080

CMD ["/app/go-execute"]
