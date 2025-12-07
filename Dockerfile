FROM golang:1.25.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application with static linking
RUN CGO_ENABLED=0 go build -tags netgo -ldflags '-s -w' -o /app/app

FROM alpine:latest

# Security: Create and switch to a non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

COPY --from=builder /app/app /app/app

EXPOSE 8080

CMD ["/app/app"]
